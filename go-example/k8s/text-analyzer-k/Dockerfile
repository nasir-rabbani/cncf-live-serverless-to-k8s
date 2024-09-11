# Builder stage: Start from the latest Go base image
FROM golang:latest AS builder

# Set the Current Working Directory inside the builder container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the builder container
COPY . .

# Build the Go app with optimizations for a smaller binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

#------
# Final stage: Use a minimal base image
FROM alpine:latest

# Set the Current Working Directory inside the final container
WORKDIR /app

# Copy the executable from the builder stage
COPY --from=builder /app/main /bin/

# Expose port 8080 to the outside world
EXPOSE 8083

# Command to run the executable
CMD ["/bin/main"]
