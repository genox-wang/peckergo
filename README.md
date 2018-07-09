# 控制台模板

### 使用方法

```bash
// 在 `{$GOPATH}/src/` 目录下下载模板工程到yourApp
git clone http://git.ti-ding.com/wangji/console-template.git yourApp

cd yourApp

// 初始化模板
go run cmd/main.go init yourApp

// mysql 数据库创建 yourApp 数据库

// 使用 https://github.com/golang/dep 下载依赖
dep ensure

// 启动后端 api 服务
go run api/main.go

// 新开命令行

cd console

// 下载前端依赖
npm install

// 运行
npm run serve

// 生产模式
npm run build
```