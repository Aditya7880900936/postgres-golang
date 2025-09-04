# ---------- Stage 1: Build ----------
    FROM golang:1.24-alpine AS builder
    WORKDIR /app
    
    # Install git
    RUN apk add --no-cache git
    
    # Copy go.mod and go.sum first (for caching)
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build the Go binary
    RUN go build -o main .
    
    # ---------- Stage 2: Runtime ----------
    FROM alpine:3.18
    WORKDIR /app
    
    # Install Postgres client (optional, for debugging)
    RUN apk add --no-cache postgresql-client
    
    # Copy the built binary
    COPY --from=builder /app/main .
    
    # Expose the app port
    EXPOSE 8080
    
    # Start the app
    CMD ["./main"]
    