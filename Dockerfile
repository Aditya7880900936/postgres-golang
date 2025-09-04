# ---------- Stage 1: Build ----------
    FROM golang:1.24-alpine AS builder
    WORKDIR /app
    
    # Install required packages
    RUN apk add --no-cache git
    
    # Copy go.mod and go.sum first for caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the project files
    COPY . .
    
    # Build Go binary
    RUN go build -o main .
    
    # ---------- Stage 2: Runtime ----------
    FROM alpine:3.20
    WORKDIR /app
    
    # Install Postgres client for debugging (optional)
    RUN apk add --no-cache postgresql-client
    
    # Copy built binary from builder
    COPY --from=builder /app/main .
    
    # Expose API port
    EXPOSE 8080
    
    # Run binary
    CMD ["./main"]
    