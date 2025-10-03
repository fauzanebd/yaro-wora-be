package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"

	appConfig "yaro-wora-be/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type StorageService struct {
	client     *s3.Client
	bucketName string
	endpoint   string
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
	imageURL := fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucketName, key)
	return imageURL, nil
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

// extractKeyFromURL extracts the object key from the full URL
func (s *StorageService) extractKeyFromURL(imageURL string) string {
	// Remove endpoint and bucket from URL to get the key
	prefix := fmt.Sprintf("%s/%s/", s.endpoint, s.bucketName)
	if strings.HasPrefix(imageURL, prefix) {
		return strings.TrimPrefix(imageURL, prefix)
	}
	return ""
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
	return fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucketName, key)
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
