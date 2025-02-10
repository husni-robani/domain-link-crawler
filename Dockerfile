FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

CMD [ "go", "run", "." ]