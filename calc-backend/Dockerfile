# syntax=docker/dockerfile:1
FROM golang:1.27-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /calc-backend

EXPOSE 9090
CMD ["/calc-backend"]
