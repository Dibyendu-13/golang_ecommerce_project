# Use a minimal base image for Go
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker's caching mechanism
COPY go.mod go.sum ./

# Download dependencies (this step is cached until go.mod or go.sum changes)
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8080

# Run the executable
CMD ["./main"]
