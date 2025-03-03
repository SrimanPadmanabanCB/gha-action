# Use the official Go base image for building
FROM golang:1.23 as builder

# Set environment for Go modules
WORKDIR /app
COPY . .

# Build the Go application (static binary for portability)
RUN go mod tidy && go build -o action-app main.go

# Use a small base image for final image
FROM alpine:latest

# Set working directory inside container
WORKDIR /app

# Copy built Go binary from builder
COPY --from=builder /app/action-app .

# Define the entrypoint to run the app
ENTRYPOINT ["/app/action-app"]