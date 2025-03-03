# Use the official Go base image for building
FROM golang:1.24.0-alpine3.21 as builder

# Set environment for Go modules
WORKDIR /app
COPY . .

# Build the Go application (static binary for portability)
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -ogo build -o action-app main.go

# Define the entrypoint to run the app
ENTRYPOINT ["/app/action-app"]