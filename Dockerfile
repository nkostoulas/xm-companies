# Start with a lightweight Golang image
FROM golang:1.23

# Set the Current Working Directory
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies first (optimizes caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the application
RUN go build -o main ./cmd

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
