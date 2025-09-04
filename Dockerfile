# ---------- Stage 1: Build ----------
    FROM golang:1.24-alpine AS builder
    WORKDIR /app
    
    # Install required tools
    RUN apk add --no-cache git bash curl build-base
    
    # Cache dependencies
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build the binary
    RUN go build -o main .
    
    # ---------- Stage 2: Runtime ----------
    FROM alpine:3.18
    WORKDIR /app
    
    # Install Postgres client for debugging
    RUN apk add --no-cache postgresql-client
    
    # Copy binary from builder
    COPY --from=builder /app/main .
    
    # Copy migration files
    COPY migrations ./migrations
    
    # Expose the app port
    EXPOSE 8080
    
    # Start the app
    CMD ["./main"]
    