FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./bin/tomata-backend

FROM alpine:latest

COPY --from=builder /app/bin/ /app/bin/

EXPOSE 80

WORKDIR /app/bin

CMD ["./tomata-backend", "0.0.0.0:80"]