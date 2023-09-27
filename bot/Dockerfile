# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /bot

# Copy the local source files into the container
COPY . .

# Build the Go application
RUN go build -o bot

# Command to run the executable
CMD ["./bot"]