# Use official Golang image as base
FROM golang:1.21

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Create directory for SQLite database
RUN mkdir -p /app/data

# Build the application
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Set environment variable for SQLite database path
ENV DB_PATH=/app/data/userdb.db

# Run the application
CMD ["./main"]
