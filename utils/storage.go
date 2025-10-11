package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	appConfig "yaro-wora-be/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
)

type StorageService struct {
	client     *s3.Client
	bucketName string
	endpoint   string // S3 API endpoint
	publicURL  string // Public URL for file access
}

var Storage *StorageService

// InitStorage initializes the Cloudflare R2 storage service
func InitStorage() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(appConfig.AppConfig.R2Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			appConfig.AppConfig.R2AccessKey,
			appConfig.AppConfig.R2SecretKey,
			"",
		)),
	)
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %v", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(appConfig.AppConfig.R2Endpoint)
	})

	Storage = &StorageService{
		client:     client,
		bucketName: appConfig.AppConfig.R2BucketName,
		endpoint:   appConfig.AppConfig.R2Endpoint,
		publicURL:  appConfig.AppConfig.R2PublicURL,
	}

	return nil
}

// UploadImage uploads an image file to R2 and returns the URL
func (s *StorageService) UploadImage(file *multipart.FileHeader, folder string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%s%s",
		strings.TrimSuffix(file.Filename, ext),
		uuid.New().String()[:8],
		ext,
	)

	// Create full path
	key := fmt.Sprintf("%s/%s", folder, filename)

	// Open file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Read file content
	fileContent, err := io.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Determine content type
	contentType := getContentType(file.Filename)

	// Upload to R2
	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        strings.NewReader(string(fileContent)),
		ContentType: aws.String(contentType),
		ACL:         "public-read", // Make publicly accessible
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Return public URL
	imageURL := fmt.Sprintf("%s/%s", s.publicURL, key)
	return imageURL, nil
}

// UploadImageWithThumbnail uploads an image and its thumbnail to R2, returns URLs and dimensions
func (s *StorageService) UploadImageWithThumbnail(file *multipart.FileHeader, folder string) (*UploadResponse, error) {
	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	baseFilename := strings.TrimSuffix(file.Filename, ext)
	uniqueID := uuid.New().String()[:8]
	filename := fmt.Sprintf("%s_%s%s", baseFilename, uniqueID, ext)

	// Create full path
	key := fmt.Sprintf("%s/%s", folder, filename)

	// Open and read file
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	fileContent, err := io.ReadAll(src)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Determine content type
	contentType := getContentType(file.Filename)

	// Check if file is SVG - skip thumbnail creation for SVG files
	isSVG := strings.ToLower(ext) == ".svg"

	var width, height int
	var thumbnailURL string

	if !isSVG {
		// Decode image to get dimensions for non-SVG files
		img, format, err := image.Decode(bytes.NewReader(fileContent))
		if err != nil {
			return nil, fmt.Errorf("failed to decode image: %v", err)
		}

		bounds := img.Bounds()
		width = bounds.Dx()
		height = bounds.Dy()

		// Generate thumbnail (10% of original size)
		thumbnailWidth := width / 10
		thumbnailHeight := height / 10
		if thumbnailWidth < 1 {
			thumbnailWidth = 1
		}
		if thumbnailHeight < 1 {
			thumbnailHeight = 1
		}

		thumbnail := imaging.Resize(img, thumbnailWidth, thumbnailHeight, imaging.Lanczos)

		// Encode thumbnail to bytes
		var thumbnailBuf bytes.Buffer
		switch format {
		case "jpeg", "jpg":
			err = imaging.Encode(&thumbnailBuf, thumbnail, imaging.JPEG)
		case "png":
			err = imaging.Encode(&thumbnailBuf, thumbnail, imaging.PNG)
		case "gif":
			err = imaging.Encode(&thumbnailBuf, thumbnail, imaging.GIF)
		default:
			err = imaging.Encode(&thumbnailBuf, thumbnail, imaging.JPEG)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to encode thumbnail: %v", err)
		}

		// Upload thumbnail
		thumbnailFilename := fmt.Sprintf("%s_%s_thumb%s", baseFilename, uniqueID, ext)
		thumbnailKey := fmt.Sprintf("%s/%s", folder, thumbnailFilename)
		_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      aws.String(s.bucketName),
			Key:         aws.String(thumbnailKey),
			Body:        bytes.NewReader(thumbnailBuf.Bytes()),
			ContentType: aws.String(contentType),
			ACL:         "public-read",
		})
		if err != nil {
			return nil, fmt.Errorf("failed to upload thumbnail: %v", err)
		}

		// Generate thumbnail URL
		thumbnailURL = fmt.Sprintf("%s/%s", s.publicURL, thumbnailKey)
	}

	// Upload original image
	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(contentType),
		ACL:         "public-read",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload original image: %v", err)
	}

	// Generate public URL
	imageURL := fmt.Sprintf("%s/%s", s.publicURL, key)

	response := &UploadResponse{
		Success:  true,
		FileURL:  imageURL,
		FileSize: file.Size,
	}

	// Only add thumbnail URL and dimensions for non-SVG files
	if !isSVG {
		response.ThumbnailURL = thumbnailURL
		response.Dimensions = &ImageDimensions{
			Width:  width,
			Height: height,
		}
	}

	return response, nil
}

// DeleteImage deletes an image from R2
func (s *StorageService) DeleteImage(imageURL string) error {
	// Extract key from URL
	key := s.extractKeyFromURL(imageURL)
	if key == "" {
		return fmt.Errorf("invalid image URL")
	}

	_, err := s.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}

// DeleteImageWithThumbnail deletes both an image and its thumbnail from R2
func (s *StorageService) DeleteImageWithThumbnail(imageURL string) error {
	// Delete the original image
	if err := s.DeleteImage(imageURL); err != nil {
		return err
	}

	// Check if the file is SVG - skip thumbnail deletion for SVG files
	key := s.extractKeyFromURL(imageURL)
	if key != "" {
		ext := filepath.Ext(key)
		// Skip thumbnail deletion for SVG files
		if strings.ToLower(ext) != ".svg" {
			// Generate thumbnail URL from original URL
			// e.g., "image_abc123.jpg" -> "image_abc123_thumb.jpg"
			baseKey := strings.TrimSuffix(key, ext)
			thumbnailKey := fmt.Sprintf("%s_thumb%s", baseKey, ext)
			thumbnailURL := fmt.Sprintf("%s/%s", s.publicURL, thumbnailKey)

			// Try to delete thumbnail (ignore error if it doesn't exist)
			_ = s.DeleteImage(thumbnailURL)
		}
	}

	return nil
}

// extractKeyFromURL extracts the object key from the full URL
func (s *StorageService) extractKeyFromURL(imageURL string) string {
	// Remove public URL prefix to get the key
	prefix := fmt.Sprintf("%s/", s.publicURL)
	if strings.HasPrefix(imageURL, prefix) {
		return strings.TrimPrefix(imageURL, prefix)
	}
	return ""
}

// IsR2URL checks if the given URL is from our R2 bucket
func (s *StorageService) IsR2URL(imageURL string) bool {
	return strings.HasPrefix(imageURL, s.publicURL)
}

// DeleteImageIfR2 deletes an image from R2 only if the URL is from our R2 bucket
func (s *StorageService) DeleteImageIfR2(imageURL string) error {
	if imageURL == "" {
		return nil // No URL to delete
	}

	if !s.IsR2URL(imageURL) {
		return nil // Not an R2 URL, skip deletion
	}

	return s.DeleteImage(imageURL)
}

// DeleteImageWithThumbnailIfR2 deletes both an image and its thumbnail from R2 only if the URL is from our R2 bucket
func (s *StorageService) DeleteImageWithThumbnailIfR2(imageURL string) error {
	if imageURL == "" {
		return nil // No URL to delete
	}

	if !s.IsR2URL(imageURL) {
		return nil // Not an R2 URL, skip deletion
	}

	return s.DeleteImageWithThumbnail(imageURL)
}

// getContentType returns the content type based on file extension
func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}

// GenerateImageURL generates a properly formatted image URL
func (s *StorageService) GenerateImageURL(key string) string {
	return fmt.Sprintf("%s/%s", s.publicURL, key)
}

// UploadResponse represents the response from image upload
type UploadResponse struct {
	Success      bool             `json:"success"`
	FileURL      string           `json:"file_url"`
	ThumbnailURL string           `json:"thumbnail_url,omitempty"`
	FileSize     int64            `json:"file_size"`
	Dimensions   *ImageDimensions `json:"dimensions,omitempty"`
}

// ImageDimensions represents image dimensions
type ImageDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// StorageAnalytics represents storage usage analytics
type StorageAnalytics struct {
	TotalSize      int64   `json:"total_size_bytes"`
	TotalSizeMB    float64 `json:"total_size_mb"`
	TotalSizeGB    float64 `json:"total_size_gb"`
	ObjectCount    int64   `json:"object_count"`
	StorageLimit   int64   `json:"storage_limit_bytes"`
	StorageLimitGB float64 `json:"storage_limit_gb"`
	UsagePercent   float64 `json:"usage_percent"`
	CanUpload      bool    `json:"can_upload"`
	RemainingBytes int64   `json:"remaining_bytes"`
	RemainingMB    float64 `json:"remaining_mb"`
}

// GetStorageAnalytics retrieves storage usage analytics for the bucket
func (s *StorageService) GetStorageAnalytics() (*StorageAnalytics, error) {
	// List all objects in the bucket to calculate total size
	var totalSize int64
	var objectCount int64

	// Use ListObjectsV2 to get all objects
	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %v", err)
		}

		for _, obj := range page.Contents {
			if obj.Size != nil {
				totalSize += *obj.Size
			}
			objectCount++
		}
	}

	// Get storage limit from config (default 1GB)
	storageLimit := int64(appConfig.AppConfig.StorageLimitGB * 1024 * 1024 * 1024)

	// Calculate analytics
	totalSizeMB := float64(totalSize) / (1024 * 1024)
	totalSizeGB := float64(totalSize) / (1024 * 1024 * 1024)
	storageLimitGB := float64(storageLimit) / (1024 * 1024 * 1024)
	usagePercent := (float64(totalSize) / float64(storageLimit)) * 100
	canUpload := totalSize < storageLimit
	remainingBytes := storageLimit - totalSize
	remainingMB := float64(remainingBytes) / (1024 * 1024)

	return &StorageAnalytics{
		TotalSize:      totalSize,
		TotalSizeMB:    totalSizeMB,
		TotalSizeGB:    totalSizeGB,
		ObjectCount:    objectCount,
		StorageLimit:   storageLimit,
		StorageLimitGB: storageLimitGB,
		UsagePercent:   usagePercent,
		CanUpload:      canUpload,
		RemainingBytes: remainingBytes,
		RemainingMB:    remainingMB,
	}, nil
}

// CheckStorageLimit checks if storage is within limits before upload
func (s *StorageService) CheckStorageLimit(fileSize int64) error {
	analytics, err := s.GetStorageAnalytics()
	if err != nil {
		return fmt.Errorf("failed to get storage analytics: %v", err)
	}

	// Check if adding this file would exceed the limit
	if analytics.TotalSize+fileSize > analytics.StorageLimit {
		return fmt.Errorf("upload would exceed storage limit. Current usage: %.2f GB, Limit: %.2f GB",
			analytics.TotalSizeGB, analytics.StorageLimitGB)
	}

	return nil
}

// DeleteImagesFromDetailSections deletes all images from DetailSections JSON that are from R2
func (s *StorageService) DeleteImagesFromDetailSections(detailSectionsJSON string) error {
	if detailSectionsJSON == "" {
		return nil // No detail sections to process
	}

	// Parse the JSON array
	var sections []map[string]interface{}
	if err := json.Unmarshal([]byte(detailSectionsJSON), &sections); err != nil {
		// If JSON parsing fails, it might not be a valid JSON array, skip deletion
		return nil
	}

	// Iterate through each section and delete images if they're from R2
	for _, section := range sections {
		if imageURL, ok := section["image_url"].(string); ok && imageURL != "" {
			// Delete the image and its thumbnail if it's from R2
			if err := s.DeleteImageWithThumbnailIfR2(imageURL); err != nil {
				// Log error but don't fail the entire operation
				fmt.Printf("Warning: Failed to delete image from detail section: %v\n", err)
			}
		}
	}

	return nil
}
