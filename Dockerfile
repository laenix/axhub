FROM golang:1.18 as builder
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE="on"
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-w -s -X main.VERSION=v0.1.0" -o app ./cmd/axhub/main.go

FROM alpine
WORKDIR /opt/
COPY --from=builder /app .
ENTRYPOINT [ "./app" ]
