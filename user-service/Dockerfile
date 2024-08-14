FROM golang:1.23

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the Go Modules
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o user-service .

# Command to run the executable
CMD ["./user-service"]