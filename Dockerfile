# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app

ADD . .
RUN go mod download
RUN go mod tidy

COPY cmd/main.go ./

RUN go build -o /main main.go

EXPOSE 8080

CMD ["/main"]