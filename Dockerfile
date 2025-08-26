FROM golang:1.24.5

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.* ./

# Download Go modules
RUN go mod download

COPY . .
RUN go install github.com/air-verse/air@latest
RUN air init
#RUN air -c .air.toml


# Copy the source code into the container

# Build the Go application
RUN go build -o main main.go

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["air", "-c", ".air.toml"]


