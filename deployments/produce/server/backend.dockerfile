FROM golang:alpine as builder

WORKDIR /app

COPY . ./

RUN go build -ldflags="-s -w" -o /app/main  ./cmd/main.go


FROM alpine as server

ENV TZ Asia/Shanghai

WORKDIR  /app

COPY --from=builder /app/configs ./
COPY --from=builder /app/main ./

EXPOSE 8080
CMD ["./main"]







