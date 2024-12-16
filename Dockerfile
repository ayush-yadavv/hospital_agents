# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary and .env file from the builder image
COPY --from=0 /app/main .
COPY --from=0 /app/.env .
COPY --from=0 /app/personalities.json .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]