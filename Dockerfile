FROM loads/alpine:3.8

LABEL maintainer="Hamster <liaolaixin@gmail.com>"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /app/main

# 添加应用可执行文件，并设置执行权限
ADD ./service   $WORKDIR/service
RUN chmod +x $WORKDIR/service

# 增加端口绑定
EXPOSE 10400


###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ["./service"]
