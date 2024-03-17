# Use the official GO Image
FROM golang:1.22-alpine

# Set the workig directory
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

RUN go mod tidy

# Download the dependencies
RUN go mod download

# Copy the source files to the container
COPY . .

# Copy the config file to the container
COPY conf/conf.yaml /app/conf/conf.yaml

# compile the application
RUN go build -o main ./src

# Expose the port 8080
EXPOSE 9090

# Run the application
CMD ["./main"]