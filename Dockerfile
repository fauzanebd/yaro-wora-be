# Build stage
FROM golang:1.21-bookworm AS builder

# Install necessary packages
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    ca-certificates \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yaro-wora-api .

# Final stage
FROM debian:bookworm-slim

# Install ca-certificates for HTTPS requests
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Create app directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/yaro-wora-api .

# Copy .env.example as reference
COPY --from=builder /app/.env .

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

# Run the application
CMD ["./yaro-wora-api"]
