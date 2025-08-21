FROM golang:1.23-alpine AS builder
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories


COPY . /app
WORKDIR /app
RUN rm -rf /app/bin/
RUN GOPROXY=https://goproxy.cn,direct go mod tidy && go build -ldflags="-s -w" -o ./bin/server ./cmd/server/main.go


FROM alpine:latest
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories


RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && mkdir -p /app/storage

WORKDIR /app
COPY --from=builder /app/bin /app

