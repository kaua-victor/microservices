FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o payment-service ./cmd/main.go

EXPOSE 3001

CMD ["./payment-service"]