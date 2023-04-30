FROM golang:1.20-buster AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o service -ldflags=-X=main.version=${VERSION} main.go

FROM loads/alpine:3.8

LABEL maintainer="Hamster <liaolaixin@gmail.com>"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /app/main
COPY --from=builder /go/src/app/service $WORKDIR/service
# 添加应用可执行文件，并设置执行权限
RUN chmod +x $WORKDIR/service

# 增加端口绑定
EXPOSE 10400

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ["./service"]






