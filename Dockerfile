# Step 1: Use Go base image to compile the code
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project directory (containing your source code)
COPY . .

# Build the Go application
RUN go build -o udpDuplicator udpDuplicator.go

# Step 2: Create a minimal final image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/udpDuplicator .

# Set the entrypoint command to start the UDP forwarder
ENTRYPOINT ["./udpDuplicator"]
