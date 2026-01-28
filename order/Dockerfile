FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o order-service ./cmd/main.go

EXPOSE 3000

CMD ["./order-service"]
