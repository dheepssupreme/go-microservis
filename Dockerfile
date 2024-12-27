# Use official Golang image as base
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
