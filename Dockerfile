# Start from the official Go image
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code (including src folder)
COPY ./src ./src

# Build the Go application
RUN go build -o /usr/local/bin/app ./src/main.go

# Start a new stage for a smaller image
FROM golang:1.21

# Copy the compiled binary from the previous stage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# Expose port 8080
EXPOSE 8080

# Set the command to run the app
CMD ["/usr/local/bin/app"]
