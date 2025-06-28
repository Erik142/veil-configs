# Use a Go image to build the application
FROM golang:1.24.3 AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker's build cache
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the server executable
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Use a minimal image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/server .

# Expose the port the server listens on (assuming default gRPC port 50051)
EXPOSE 50051

# Command to run the executable
ENTRYPOINT ["./server"]
