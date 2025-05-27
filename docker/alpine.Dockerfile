# 构建阶段
FROM golang:1.24 AS builder

RUN apt install -y git

WORKDIR /app

#RUN #sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /bin

COPY ../go.mod go.sum ./
RUN go mod download

COPY Taskfile.yml .

COPY .. .

RUN go tool task build-linux-docker

# 运行阶段
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/proxite /root/proxite

EXPOSE 9876

CMD ["/root/proxite"]
