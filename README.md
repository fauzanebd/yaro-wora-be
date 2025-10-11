# Yaro Wora Tourism Website - Backend API

A complete Golang backend API for the Yaro Wora tourism website built with Fiber framework, PostgreSQL database, and Cloudflare R2 for image storage.

## Project Structure

```
yaro-wora-be/
├── config/          # Configuration management
├── models/          # Database models
├── handlers/        # API route handlers
├── middleware/      # Custom middleware
├── utils/           # Utility functions
├── migrations/      # Database migrations
├── main.go          # Application entry point
├── go.mod           # Go module dependencies
└── README.md        # This file
```

## Prerequisites

- Go 1.19 or higher
- PostgreSQL 12 or higher
- Cloudflare R2 account (for image storage)

## Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd yaro-wora-be
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env
   ```

   Edit `.env` file with your configuration:

   ```env
   # Database Configuration
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=yaro_wora

   # Cloudflare R2 Configuration
   R2_ACCESS_KEY=your_r2_access_key
   R2_SECRET_KEY=your_r2_secret_key
   R2_BUCKET_NAME=yaro-wora-images
   R2_ENDPOINT=https://your-account-id.r2.cloudflarestorage.com

   # Server Configuration
   PORT=3000
   JWT_SECRET=your-super-secret-jwt-key

   # Admin Authentication
   ADMIN_USERNAME=admin
   ADMIN_PASSWORD=admin123
   ```

4. **Set up PostgreSQL database**

   ```bash
   createdb yaro_wora
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```
