# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the current folder contents to the working directory
COPY . .

# Download dependencies and build the application
RUN go mod tidy && go build -o cnc .

# Stage 2: Create the runtime environment
FROM ubuntu:22.04

# Set environment variables for non-interactive installations
ENV DEBIAN_FRONTEND=noninteractive

# Copy the application binary from the builder stage
COPY --from=builder /app/cnc /cnc

CMD /cnc

# Expose the necessary ports
EXPOSE 23
