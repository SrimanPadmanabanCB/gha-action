# Use the official Go base image for building
FROM golang:1.24 as builder

# Set environment for Go modules
WORKDIR /app
COPY . .

# Build the Go application (static binary for portability)
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -o action-app main.go

FROM gcr.io/distroless/base:latest

COPY --from=builder /app/action-app /app/action-app

# Define the entrypoint to run the app
ENTRYPOINT ["/app/action-app"]