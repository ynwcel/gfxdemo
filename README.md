### 1、安装
#### 1.1 下载源码
```
git clone https://github.com/ynwcel/gfxdemo
```
#### 1.2 重命名包表（可选）
```
go run hack/hack.go --setgomod=mygoweb
```

### 2、编译
#### 2.1 自动设置版本及日期
```
go run hack/hack.go --build --build.version=0.0.1
```

