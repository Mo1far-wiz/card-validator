# syntax=docker/dockerfile:1

FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/main /app/cmd/api/

EXPOSE 8080

CMD ["/app/bin/main"]

