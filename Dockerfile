# Stage 1: Builder
# Use a specific Go version for reproducibility
FROM golang:1.24.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
# This step is only invalidated if go.mod or go.sum changes
COPY go.mod go.sum ./

# Download dependencies
# This step is only invalidated if go.mod or go.sum changes and new dependencies are needed
RUN go mod download

# Copy the rest of the application source code
# This step is invalidated if any source file changes
COPY . .

# Build the application
# CGO_ENABLED=0 is used to create a statically linked binary, making the final image smaller
# -o main specifies the output filename
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix netgo -o main main.go

# Stage 2: Final image
# Use a minimal base image like alpine for a small final image
FROM alpine:latest

# Install ca-certificates to ensure SSL connections work
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
ENTRYPOINT ["./main"]
