# 构建阶段
FROM golang:1.24 AS builder

RUN apt install -y git

WORKDIR /app

RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

COPY ../go.mod go.sum ./
RUN go mod download

COPY .. .

RUN task build-linux

# 运行阶段
FROM debian:stable

RUN apt update && apt install -y ca-certificates

WORKDIR /root/

COPY --from=builder /app/proxite .

EXPOSE 9876

CMD ["./proxite"]
