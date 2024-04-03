# assignment_db




# Start with the official Golang image
FROM golang:latest AS builder

# Setting the Current Working Directory inside the container
WORKDIR /app

# Copying the Go mod and sum files
COPY go.mod go.sum ./

# Downloading Go modules. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copying the source code from the current directory to the Working Directory inside the container
COPY . .

# Building the Go app
RUN go build -o main .

# Intermediate stage to create a smaller final image
FROM golang:alpine

# Setting the Current Working Directory inside the container
WORKDIR /app

# Copying the binary from the builder stage to the smaller image
COPY --from=builder /app/main .

# Installing MongoDB tools
RUN apk add --no-cache mongodb-tools

# Exposing port 8080 for the Go application
EXPOSE 8080

#running the Go application
CMD ["./main"]
