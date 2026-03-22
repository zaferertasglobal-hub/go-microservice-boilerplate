# Use the official Golang image as a build stage  
FROM golang:1.20 as builder  

# Set the Current Working Directory inside the container  
WORKDIR /app  

# Copy go modules manifests  
COPY go.mod .  
COPY go.sum .  

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed  
runs "go mod download"  

# Copy the source code into the container  
COPY . .  

# Build the Go app  
RUN CGO_ENABLED=0 GOOS=linux go build -o main .  

# Start a new stage from scratch  
FROM alpine:latest  

# Set the Current Working Directory inside the container  
WORKDIR /root/  

# Copy the Pre-built binary file from the previous stage  
COPY --from=builder /app/main .  

# Command to run the executable  
CMD ["./main"]