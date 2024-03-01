FROM golang:1.22.0-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add curl
RUN git config --global --unset http.proxy
RUN git config --global --unset https.proxy
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64.tar.gz ./migrate
COPY serverConfig.json .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./migrations

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]