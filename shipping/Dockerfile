FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o shipping-service ./cmd/main.go

EXPOSE 3002

CMD ["./shipping-service"]
