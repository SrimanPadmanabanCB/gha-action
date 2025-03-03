# Use the official Go base image for building
FROM golang:1.24 as builder

# Set environment for Go modules
WORKDIR /app
COPY . .

# Build the Go application (static binary for portability)
RUN go mod tidy && go build -a -tags "$BUILD_TAGS" -ldflags '-w -extldflags \"-static\"' -o action-app main.go

FROM alpine:latest

# Set working directory inside container
WORKDIR /app
COPY --from=builder /app/action-app .

# Define the entrypoint to run the app
ENTRYPOINT ["/app/action-app"]