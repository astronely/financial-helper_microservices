FROM golang:1.24-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mode \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o main cmd/main.go

EXPOSE 50052

CMD ["./main"]