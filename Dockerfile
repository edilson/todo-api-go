FROM golang:1.25.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todo-api

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo-api .

EXPOSE 8080

CMD [ "./todo-api" ]