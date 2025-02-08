# Use the official Go image from Docker Hub
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of your application
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

# Expose port 8080 for the application
EXPOSE 8080

# Set the default command to run the application
CMD ["./finance"]
