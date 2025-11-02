# Multi-stage Dockerfile for a Go app. Adjust BUILD_PATH if your main package lives elsewhere.
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src

ARG BUILD_PATH=.
ARG BINARY=meditation-app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/${BINARY} ${BUILD_PATH}

# Final minimal runtime image
FROM alpine:3.18 AS runtime
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /out/meditation-app /app/meditation-app

# non-root user
RUN addgroup -S app && adduser -S -G app app
USER app

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app/meditation-app"]