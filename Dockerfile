# Build stage
FROM golang:1.22 as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o proxy-server main.go

# Final stage
FROM gcr.io/distroless/static

WORKDIR /

# Copy binary from builder
COPY --from=builder /app/proxy-server .

# Run the Go app
ENTRYPOINT ["/proxy-server"]grafana
