
# grpc-gateway 示例项目
- 一个grpc-server服务,对外提供grpc端口
- 一个gateway-server服务,对外提供http端口

## 依赖 buf 来统一 proto 插件, buf 常用命令
```shell
buf --version
buf config init
buf build
buf generate
buf lint
buf breaking --against ".git#subdir=./proto"
```
## 详细使用方法见Makefile