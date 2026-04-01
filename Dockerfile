FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go mod download && CGO_ENABLED=0 go build -o brander ./cmd/brander/

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/brander .
ENV PORT=9180 DATA_DIR=/data
EXPOSE 9180
CMD ["./brander"]
