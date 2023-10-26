FROM golang:1.20

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go build -o main .

EXPOSE 8080