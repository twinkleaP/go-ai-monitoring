# Dockerfile.go
FROM golang:1.24 AS builder


WORKDIR /app

# Copy go.mod first
COPY go.mod ./
# Copy go.sum only if it exists
COPY go.sum* ./


# Copy Go source code
COPY . .

EXPOSE 8080

# Run Go service
CMD ["go", "run", "main.go"]
