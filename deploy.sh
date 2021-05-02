#!/bin/bash

# 确保脚本抛出遇到的错误
set -e

# 编译前端代码

echo "start build console..."

cd console
npm install
npm run build

echo "start copy console to remote ..."

# 删除服务器已有目录
ssh root@47.103.147.164 "rm -rf /www/peckergo"

# 静态文件拷贝到服务器
scp -r dist root@47.103.147.164:/www/peckergo

cd ../

echo "start build api..."

# 编译后端代码
go mod download

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o peckergo ./api

chmod +x peckergo

echo "start copy api to remote ..."

# 可执行文件拷贝到服务器
scp peckergo root@47.103.147.164:/root/app/peckergo/peckergo_tmp
scp api/config/config.yml.prod root@47.103.147.164:/root/app/peckergo/api/config/
scp docker-compose.yml root@47.103.147.164:/root/app/peckergo/
scp Dockerfile root@47.103.147.164:/root/app/peckergo/
scp zoneinfo.zip root@47.103.147.164:/root/app/peckergo/
