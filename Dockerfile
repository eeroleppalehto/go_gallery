# Use Go 1.23 bookworm as base image
FROM golang:1.23-alpine AS base

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o gollery cmd/main.go

# Document the port that may need to be published
EXPOSE 8081

# Start the application
CMD ["/build/gollery"]
