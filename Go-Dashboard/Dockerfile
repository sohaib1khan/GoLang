# Stage 1: Build the Go application
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests first (for caching dependencies)
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the entire application code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Create a minimal runtime container
FROM alpine:latest

# Install necessary tools (e.g., CA certificates for HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory in the runtime container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main .

# Copy required static files and views
COPY static /app/static
COPY views /app/views
COPY data.json /app/data.json

# Ensure the binary is executable
RUN chmod +x ./main

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./main"]
