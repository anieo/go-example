FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download

RUN go build -o app

FROM alpine:3.14

COPY --from=0 /app/app /app/app

WORKDIR /app

EXPOSE 8080

CMD ["./app"]