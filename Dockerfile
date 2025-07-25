# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Set Go proxy and module settings
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV GO111MODULE=on

# Add git and bash (optional) for private repositories or debugging
RUN apk add --no-cache git

# Copy go.mod and go.sum first
COPY go.mod ./
COPY go.sum ./

# Download dependencies (including postgres driver)
RUN go mod tidy

# Copy environment file (optional)
COPY .env ./

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN go build -o main .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary and .env file from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port (adjust to your app)
EXPOSE 8081

# Start the service
CMD ["./main"]
