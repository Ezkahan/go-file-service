# --- Stage 1: Build the Go binary ---
FROM golang:1.25.3-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src

ARG BUILD_PATH=./cmd/app
ARG BINARY=meditation-app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/${BINARY} ${BUILD_PATH}

# --- Stage 2: Runtime image ---
FROM alpine:3.18 AS runtime
RUN apk add --no-cache ca-certificates bash curl postgresql-client

WORKDIR /app

# Install golang-migrate CLI
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz \
    | tar -xz -C /usr/local/bin && chmod +x /usr/local/bin/migrate

# Copy binary and migrations
COPY --from=builder /out/meditation-app /app/meditation-app
COPY internal/db/migrations internal/db/migrations
COPY .env .env

# --- create entrypoint script while still root ---
COPY <<'EOF' /app/entrypoint.sh
#!/bin/bash
set -e

echo "⏳ Waiting for PostgreSQL to be ready..."
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" >/dev/null 2>&1; do
  sleep 2
done

echo "✅ Database is up. Running migrations..."
migrate -path internal/db/migrations -database "$DATABASE_URL" up

echo "🚀 Starting Go app..."
exec /app/meditation-app
EOF

# make script executable (we're still root here)
RUN chmod +x /app/entrypoint.sh

# --- create non-root user *after* script is set up ---
RUN addgroup -S app && adduser -S -G app app
USER app

EXPOSE 8080
ENTRYPOINT ["/app/entrypoint.sh"]
