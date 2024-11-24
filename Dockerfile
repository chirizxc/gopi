# Use an official Go image as a build stage
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory (excluding ignored files like those in .dockerignore)
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Use Alpine as a minimal runtime image
FROM alpine:3.18

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the config file to the container
COPY config.toml .

# Install curl and other necessary tools
RUN apk add --no-cache libc6-compat curl

# Expose the port on which the app will run
EXPOSE 8000

# Command to run the application with delay
CMD ["sh", "-c", "sleep 10 && ./main"]
