FROM golang:1.23-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gollery cmd/main.go

EXPOSE 8081

CMD ["/build/gollery"]
