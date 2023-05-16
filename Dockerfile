FROM golang:1.20-buster AS builder

# 设置构建参数的默认值
ARG GIT_TAG="unknown"
ARG GIT_COMMIT_LOG="unknown"
ARG BUILD_TIME="unknown"
ARG BUILD_GO_VERSION="github@action golang:1.20-buster"


WORKDIR /go/src/app
COPY . .

# 打印构建参数
RUN echo "GIT_TAG=${GIT_TAG}"
RUN echo "GIT_COMMIT_LOG=${GIT_COMMIT_LOG}"
RUN echo "BUILD_TIME=${BUILD_TIME}"
RUN echo "BUILD_GO_VERSION=${BUILD_GO_VERSION}"


# 设置 LDFlags 变量
ENV LDFLAGS=" \
    -X 'push/utility/bin_utils.GitTag=${GIT_TAG}' \
    -X 'push/utility/bin_utils.GitCommitLog=${GIT_COMMIT_LOG}' \
    -X 'push/utility/bin_utils.BuildTime=${BUILD_TIME}' \
    -X 'push/utility/bin_utils.BuildGoVersion=${BUILD_GO_VERSION}' \
"

RUN CGO_ENABLED=0 go build -o service -ldflags "$LDFLAGS" main.go

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




