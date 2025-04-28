FROM golang:1.24-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod tidy
RUN go mod download
COPY . .

RUN go build -o main cmd/main.go

# Expose the port the service will run on
EXPOSE 8080 9090

# Command to run the application
CMD ["./main"]