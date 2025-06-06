# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git ca-certificates tzdata
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ephemeral-port-exporter ./cmd/ephemeral-port-exporter

# Final stage
FROM alpine:latest
RUN apk --no-cache add iproute2 ca-certificates bash coreutils
RUN addgroup -g 1001 appuser && adduser -D -u 1001 -G appuser appuser
WORKDIR /app
COPY --from=builder /app/ephemeral-port-exporter .
RUN chown -R appuser:appuser /app
USER appuser
EXPOSE 2112

CMD ["./ephemeral-port-exporter"]
