# =======================
# Stage 1: Build the binary
# =======================
FROM golang:1.22-alpine AS build

# Install git and build tools (if any dependency needs C)
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary for the host architecture
RUN go build -o taskly-cli ./cmd/taskly

# =======================
# Stage 2: Production image
# =======================
FROM alpine:3.18 AS production

# Install necessary libraries
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/taskly-cli .

# Expose port if REST API is used
EXPOSE 3000

# Command to run
CMD ["./taskly-cli"]
