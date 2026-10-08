# --- Stage 1: Build the Go Binary ---
FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build a statically linked binary for Linux
RUN go build -o auth-service ./cmd/server

# --- Stage 2: Final Lightweight Runtime ---
FROM alpine:latest

# Install ca-certificates in case your Go app makes HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/auth-service .

# Expose the port your Go app listens on
EXPOSE 8080

# Run the binary
CMD ["./auth-service"]
