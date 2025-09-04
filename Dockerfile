# ======================
# Stage 1: Build the Go App
# ======================
FROM golang:1.24-alpine AS builder

# Install required packages
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Cache Go modules first
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# ======================
# Stage 2: Run the Go App
# ======================
FROM alpine:3.19

# Install CA certificates for secure DB connections
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/main .

# Copy .env file (optional, better to pass via docker-compose)
COPY .env .

# Expose port
EXPOSE 8080

# Start the app
CMD ["./main"]
