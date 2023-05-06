FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /app

COPY . ./

RUN go build -ldflags="-s -w" -o /app/main  ./cmd/main.go


FROM alpine as server

ENV TZ Asia/Shanghai

ENV chatgpt-web-log "/app/configs/log_produce.yaml"
ENV chatgpt-web-databse "/app/configs/databse_produce.yaml"
WORKDIR  /app

COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/main ./

EXPOSE 8080
CMD ["./main"]







