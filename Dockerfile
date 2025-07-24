# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Set Go proxy and module settings
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV GO111MODULE=on

# Add git for private repositories (if needed)
RUN apk add --no-cache git

# Copy go.mod firstdocker ps

COPY go.mod ./
COPY .env ./

# Copy go.sum only if it exists
COPY go.su[m] ./

RUN go mod download

COPY . .
RUN go build -o main .

# Runtime stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary and .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port (adjust based on your app)
EXPOSE 8081

# Run the binary
CMD ["./main"]
