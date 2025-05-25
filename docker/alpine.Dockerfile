# 构建阶段
FROM golang:1.24 AS builder

RUN apt install -y git

WORKDIR /app

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o proxite .

# 运行阶段
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/proxite .

EXPOSE 9876

CMD ["./proxite"]
