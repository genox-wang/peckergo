# PeckerGo

![初始化项目](https://github.com/wilfordw/peckergo/blob/master/pecker_init_example.gif?raw=true)

![添加 model](https://github.com/wilfordw/peckergo/blob/master/pecker_model_example.gif?raw=true)

### 简介

- 前端技术 [Vue](https://github.com/vuejs/vue) + [iView](https://github.com/iview/iview) + [iView-Admin](https://github.com/iview/iview-admin) + [Webpack](https://github.com/webpack/webpack) 
- 后端技术 [Gin](https://github.com/gin-gonic/gin) + [Gorm](https://github.com/gin-gonic/gin) + [Viper](https://github.com/spf13/viper)

基于以上技术开发的一键生成部署的，Admin 项目模板

### 特点

- 提供命令行工具 [pecker](http://git.ti-ding.com/wangji/pecker) 初始化项目，自动添加 model 代码
- 后端支持跨域
- 前后端登陆对接图片验证码支持
- 用户认证基于 JWT

### 准备工作

安装 [packer](http://git.ti-ding.com/wangji/pecker)

### 使用方法

##### 初始化

```bash
// 在 `{$GOPATH}/src/` 目录下下载模板工程到yourApp
git clone http://git.ti-ding.com/wangji/peckergo.git yourApp

cd yourApp

// 初始化模板
pecker init

```

##### 运行后端代码 (需要 golang 10 + dep)

```
// mysql 数据库创建 yourApp 数据库

// 使用 https://github.com/golang/dep 下载依赖
dep ensure -v

cd api

// 启动后端 api 服务
go run main.go

```

##### 运行前端代码 (需要 npm 6 + node 10)

```
cd console

// 下载前端依赖
npm install

// 运行
npm run serve

```

##### 项目部署 (需要 golang 10 + dep + npm 6 + node 10 + docker + docekr-compose)

```
// 项目克隆到服务器 src 文件下 

cp yourApp

./build.sh
```

### 实践项目

[trans-trip-admin](http://git.ti-ding.com/hastrans/trans-trip-admin)