# Stage 1: Build the bot
FROM ubuntu:20.04 AS build

# Set environment variables for non-interactive apt
ENV DEBIAN_FRONTEND=noninteractive

# Install necessary tools for building the bot
RUN apt-get update && apt-get install -y \
    build-essential \
    gcc \
    make \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /build

# Copy the source code into the container
COPY . /build

# Compile the source code
RUN make

# Stage 2: Create the runtime image
FROM ubuntu:20.04

# Set environment variables for non-interactive apt
ENV DEBIAN_FRONTEND=noninteractive

# Install any runtime dependencies (minimal)
RUN apt-get update && apt-get install -y \
    libstdc++6 \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /build/bot /app/bot

# Expose any required ports (customize as needed)
# EXPOSE 23 80 443

# Command to run the bot
CMD ["./bot"]
