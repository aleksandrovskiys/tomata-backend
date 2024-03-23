FROM golang:1.21.1-alpine as builder

WORKDIR /app

# needed to work with cgo
ENV CGO_ENABLED=1
run apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/tomata-backend

FROM alpine:latest

COPY --from=builder /app/bin/ /app

EXPOSE 80

WORKDIR /app

CMD ["./tomata-backend", "0.0.0.0:80"]
