# card-validator

This is api service for Credit Card number and Expiration date validation.

## How to Run

```cmd
git clone https://github.com/Mo1far-wiz/card-validator.git
```
```cmd
cd card-validator/
```
For Docker:
```cmd
docker-compose up --build
```
For CMD:
```cmd
go mod tidy
```
```cmd
go run ./cmd/api
```
To run tests:
```cmd
go test ./tests -v
```

## External packages used

- [Chi](https://pkg.go.dev/github.com/go-chi/chi/v5): router for building Go HTTP services
- [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10@v10.23.0#section-readme): validation for structs and individual fields based on tags
- [Godotenv](https://pkg.go.dev/github.com/joho/godotenv): loads env vars from a .env file

## Dockerfile and Docker-compose

Project contains Dockerfile and Docker-compose files for easier deployment (also it was requirement 😁).

For me size of the created image was approximately ~250 mb.

! Dockerfile exposes port 8080.

## File structure
```
.
├── Dockerfile
├── README.md
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api.go
│   │   └── handlers
│   │       ├── errors.go
│   │       ├── health.go
│   │       └── validate-card.go
│   ├── config
│   │   └── env.go
│   ├── domain
│   │   ├── models
│   │   │   └── card.go
│   │   └── validator
│   │       ├── card-validator-interface.go
│   │       ├── card-validator.go
│   │       └── errors.go
│   └── utils
│       ├── json
│       │   └── json.go
│       └── luhn
│           └── algorithm.go
└── tests
    ├── card-validation_test.go
    ├── luhn-alg_test.go
    └── mock
        └── validator.go
```