FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o ./main ./cmd/main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/deployments /app/deployments

EXPOSE 8080

CMD ["/app/main", "-config", "/app/deployments/local.yaml"]