package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// Cloudflare R2
	R2AccessKey  string
	R2SecretKey  string
	R2BucketName string
	R2Endpoint   string // S3 API endpoint for operations
	R2PublicURL  string // Public URL for accessing files (r2.dev or custom domain)
	R2Region     string

	// Server
	Port      string
	JWTSecret string

	// Admin Auth
	AdminUsername string
	AdminPassword string

	// App
	AppEnv     string
	APIVersion string

	// Upload
	MaxFileUploadSize int     // in bytes
	StorageLimitGB    float64 // in GB
}

var AppConfig *Config
var DB *gorm.DB

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = &Config{
		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "yaro_wora"),

		// Cloudflare R2
		R2AccessKey:  getEnv("R2_ACCESS_KEY", ""),
		R2SecretKey:  getEnv("R2_SECRET_KEY", ""),
		R2BucketName: getEnv("R2_BUCKET_NAME", "yaro-wora-images"),
		R2Endpoint:   getEnv("R2_ENDPOINT", ""),
		R2PublicURL:  getEnv("R2_PUBLIC_URL", ""),
		R2Region:     getEnv("R2_REGION", "auto"),

		// Server
		Port:      getEnv("PORT", "3000"),
		JWTSecret: getEnv("JWT_SECRET", "default-secret-change-this"),

		// Admin Auth
		AdminUsername: getEnv("ADMIN_USERNAME", "admin"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123"),

		// App
		AppEnv:     getEnv("APP_ENV", "development"),
		APIVersion: getEnv("API_VERSION", "v1"),

		// Upload - Default 4MB (4 * 1024 * 1024 = 4194304 bytes)
		MaxFileUploadSize: getEnvAsInt("MAX_FILE_UPLOAD_SIZE_IN_BYTES", 4194304),
		StorageLimitGB:    getEnvAsFloat("STORAGE_LIMIT_GB", 1.0), // Default 1GB
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue
		}
	}
	return defaultValue
}

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}
