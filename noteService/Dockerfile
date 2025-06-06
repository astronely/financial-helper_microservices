FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -v -p 4 -ldflags='-w -s' \
    -o main cmd/main.go

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS and create non-root user
RUN apk --no-cache add ca-certificates \
    && addgroup -g 1001 -S appgroup \
    && adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy binary and set ownership
COPY --from=builder --chown=appuser:appgroup /build/main .
COPY --from=builder /build/local.env /app/local.env
COPY --from=builder /build/prod.env /app/prod.env

# Switch to non-root user
USER appuser

EXPOSE 50054


CMD ["./main"]