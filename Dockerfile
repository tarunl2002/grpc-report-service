# Use official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Install dependencies
RUN go mod download

# Build the Go binary
RUN go build -o grpc-service main.go

# Expose gRPC port
EXPOSE 50051

# Run the binary
CMD ["./grpc-service"]
