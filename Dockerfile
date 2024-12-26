FROM golang:alpine AS builder

COPY . .
RUN go build -o /app ./cmd


FROM alpine:latest AS run

COPY --from=builder /app /app
RUN chmod +x /app
ENTRYPOINT [ "/app" ]