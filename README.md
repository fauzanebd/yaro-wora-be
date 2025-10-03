# Yaro Wora Tourism Website - Backend API

A complete Golang backend API for the Yaro Wora tourism website built with Fiber framework, PostgreSQL database, and Cloudflare R2 for image storage.

## Features

- 🚀 **High Performance**: Built with Go Fiber framework
- 🗄️ **PostgreSQL Database**: Full-featured database with GORM ORM
- ☁️ **Cloudflare R2 Storage**: Seamless image uploads and storage
- 🔐 **Simple Authentication**: JWT-based authentication with configurable basic auth
- 📱 **Multi-language Support**: English and Indonesian localization
- 📊 **Complete API**: All endpoints for managing tourism content
- 🛡️ **Secure**: Input validation, error handling, and security middleware
- 📝 **Well Documented**: Clean code structure and comprehensive API docs

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

The server will start on `http://localhost:3000`

## API Endpoints

### Public Endpoints

- `GET /v1/carousel` - Get carousel slides
- `GET /v1/attractions` - Get attractions
- `GET /v1/pricing` - Get entrance pricing
- `GET /v1/profile` - Get village profile
- `GET /v1/destinations` - Get destinations
- `GET /v1/gallery` - Get gallery images
- `GET /v1/regulations` - Get regulations
- `GET /v1/facilities` - Get facilities
- `GET /v1/news` - Get news articles
- `POST /v1/contact` - Submit contact form
- `POST /v1/facilities/:id/book` - Book a facility

### Authentication

- `POST /v1/auth/login` - Simple login
- `POST /v1/auth/jwt-login` - JWT-based login

### Admin Endpoints (Protected)

All admin endpoints require authentication via `Authorization: Bearer <token>` header.

- `POST /v1/admin/carousel` - Create carousel slide
- `PUT /v1/admin/carousel/:id` - Update carousel slide
- `DELETE /v1/admin/carousel/:id` - Delete carousel slide
- `POST /v1/admin/attractions` - Create attraction
- `PUT /v1/admin/attractions/:id` - Update attraction
- `DELETE /v1/admin/attractions/:id` - Delete attraction
- `PUT /v1/admin/pricing` - Update pricing
- `PUT /v1/admin/profile` - Update profile
- `GET /v1/admin/contacts` - Get contact submissions
- `GET /v1/admin/bookings` - Get bookings
- `POST /v1/admin/content/upload` - Upload files

## Database Models

The application includes comprehensive models for:

- **User**: Admin user management
- **Carousel**: Hero section slides
- **Attraction**: Featured attractions
- **Pricing**: Entrance fee pricing
- **Profile**: Village profile information
- **Destination**: Tourism destinations
- **Gallery**: Image gallery with categories
- **Regulation**: Rules and guidelines
- **Facility**: Accommodations and experiences
- **News**: News articles and categories
- **Contact**: Contact form submissions
- **Booking**: Facility bookings

## Authentication

The API supports two authentication modes:

1. **Simple Authentication**: Basic username/password for the simple tier
2. **JWT Authentication**: Token-based authentication for advanced features

### Simple Auth Example

```bash
curl -X POST http://localhost:3000/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin123"}'
```

### JWT Auth Example

```bash
curl -X GET http://localhost:3000/v1/admin/profile \
  -H "Authorization: Bearer your-jwt-token"
```

## Image Upload

Images are stored in Cloudflare R2 (S3-compatible storage). The application:

1. Accepts multipart form uploads
2. Generates unique filenames
3. Uploads to R2 bucket
4. Returns public URLs for database storage

### Upload Example

```bash
curl -X POST http://localhost:3000/v1/admin/content/upload \
  -H "Authorization: Bearer your-jwt-token" \
  -F "file=@image.jpg" \
  -F "folder=carousel"
```

## Development

### Run with hot reload

```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Database Migration

The application automatically runs migrations on startup. To reset the database:

```bash
# Drop and recreate database
dropdb yaro_wora
createdb yaro_wora

# Restart the application to run migrations
go run main.go
```

### Adding New Models

1. Create model in `models/` directory
2. Add to `models/models.go` AutoMigrate function
3. Create handlers in `handlers/` directory
4. Add routes in `main.go`

## Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...
```

## Deployment

### Using Docker (recommended)

```dockerfile
FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
```

### Environment Variables for Production

Make sure to set secure values for:

- `JWT_SECRET`: Use a long, random string
- `ADMIN_PASSWORD`: Use a strong password
- Database credentials
- R2 credentials with appropriate permissions

## API Documentation

Full API documentation is available in:

- `API_DOCUMENTATION.md` - Public endpoints
- `ADMIN_API_DOCUMENTATION.md` - Admin endpoints

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is proprietary software for Yaro Wora Tourism.

## Support

For support, email your-email@example.com or create an issue in the repository.

---
