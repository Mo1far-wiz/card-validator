# syntax=docker/dockerfile:1

# first image (stage) with everything
FROM golang:1.23-alpine AS building

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/main /app/cmd/api/

# second image optimized in size
FROM golang:1.23-alpine AS runtime

WORKDIR /app

COPY --from=building /app/bin/main .
# COPY --from=building /app/.env .

EXPOSE 8080

ENTRYPOINT ["./main"]

