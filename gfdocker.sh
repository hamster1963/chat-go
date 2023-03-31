#!/bin/bash

# 初始化
gf gen dao
gf gen service
# 可将任意文件打包后随着可执行文件一同发布，就是发布只需要一个运行文件就可以
echo -e "\n" | gf pack manifest/config ./internal/packed/packed.go -n packed
# 打包文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service main.go

# 获取版本号
version=$(yq '.gfcli.build.version' ./hack/config.yaml)
echo "version: $version"
# 构建docker镜像
dockerHubName=hamster1963/push-go
docker build -t $dockerHubName:"$version" .
docker tag $dockerHubName:"$version" $dockerHubName:latest
# 推送镜像
docker push $dockerHubName:"$version"
docker push $dockerHubName:latest

