# Build executable stage
FROM golang:1.21.5-alpine3.18 as builder
WORKDIR /build
COPY . .
RUN go build -o main ./cmd/main.go
# Build final image
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /build/main ./
ENTRYPOINT ["/app/main"]