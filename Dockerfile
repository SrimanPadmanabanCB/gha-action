# Use the official Go base image for building
FROM golang:1.24.0-alpine3.21

# Set environment for Go modules
WORKDIR /app
COPY . .

# Build the Go application (static binary for portability)
RUN go mod tidy
RUN go build -a -tags "$BUILD_TAGS" -ldflags '-w -extldflags \"-static\"' -o action-app main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w -extldflags "-static"' -ogo build -o action-app main.go

# Define the entrypoint to run the app
ENTRYPOINT ["/app/action-app"]