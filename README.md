# 说明文档

### 准备

1. [安装Go](https://www.runoob.com/go/go-environment.html)
2. [安装glide](https://github.com/Masterminds/glide)
3. 安装node v10+
4. [安装Docker](https://docs.docker.com/install/)
5. [安装Docker Compose](https://docs.docker.com/compose/install/)
6. 安装mysql8，建库 peckergo
7. 安装redis

### 本地运行

假设本地已搭建好mysql8[3306],redis[6379]

> 如果端口不一样请在 `api/config/config.yml` 里配置

启动 API

```shell
glide install  # 有些库可能需要全局代理

# 开关代理方法
# alias setproxy="export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087;"
# alias unsetproxy="unset http_proxy;unset https_proxy"

cd api

go run main.go 
# 启动默认 8000 端口
```
启动前端页面

```shell

cd console

npm install

npm run sever
```

打开 http://localhost:8080/ 就能访问管理页面，初始账号密码 admin admin

> 身份 IP 过滤支持配置纯真IP服务  https://github.com/wilfordw/qqwry，参考里面的 Docker 部署
> ip 服务配置在 api/config/config.yml -> ip 项

### 生产环境部署

创建 /www 目录用于部署静态页面

配置 `config.yml.prod` 里相关参数

```shell
./build.sh
```
nginx 配置

```
server {
  listen            80;
  server_name       adv.api.mcbox.cn;

  access_log /var/log/nginx/adv.api.access.log main;
  #access_log off;
  error_log /var/log/nginx/adv.api.error.log;


  location / {
     try_files /_not_exists_ @backend;
  }

  location @backend {
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header Host            $http_host;

    proxy_pass http://127.0.0.1:8901;
  }
}
```

```
server {
  listen            80;
  server_name       adv.console.mcbox.cn;

  root /www/peckergo;
  index index.html;

  gzip on;
  gzip_min_length 1k;
  gzip_comp_level 2;
  gzip_types text/plain application/javascript application/x-javascript text/css application/xml text/javascript application/x-httpd-php image/jpeg image/gif image/png;
  gzip_vary on;
  gzip_disable "MSIE [1-6]\.";

  location / {
      try_files $uri $uri/ = 404;
  }
}
```

### 常用 docker 命令

```shell
docker ps #插件运行中 docker 容器
docker stop 容器名 #停用容器
docker start 容器名 #启动容器
docker restart 容器名 #重启容器
docker logs -f 容器名 #查看容器日志
```
### 关于广告扩展

接入的 API 逻辑都在 `api/logic` 文件夹下。目前只对接了移云和众盟，如有新增广告主，在里面添加转换逻辑即可