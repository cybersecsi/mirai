# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the builder container
WORKDIR /app

# Copy the current folder contents to the working directory
COPY . .

# Download dependencies and build the application
RUN go mod tidy && go build -o cnc .

# Stage 2: Create the runtime environment with MySQL
FROM ubuntu:22.04

# Set environment variables for non-interactive installations
ENV DEBIAN_FRONTEND=noninteractive

# Install MySQL server and client, and configure it
RUN apt-get -y update \
    && echo "#!/bin/sh\nexit 0" > /usr/sbin/policy-rc.d \
    && apt-get install -y debconf-utils \
    && echo mysql-server mysql-server/root_password password root | debconf-set-selections \
    && echo mysql-server mysql-server/root_password_again password root | debconf-set-selections \
    && apt-get install -y mysql-server -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" --fix-missing \
    && apt-get install -y mysql-client

# Copy the application binary from the builder stage
COPY --from=builder /app/cnc /cnc

# Optionally copy an SQL file for database initialization
# Assuming your SQL file is in the current directory named "db.sql"
COPY db.sql /db.sql

# Set the default command to start MySQL and run the app
CMD service mysql start \
    && mysql -u root -p"root" < /db.sql \
    && /cnc

# Expose the necessary ports
EXPOSE 23
