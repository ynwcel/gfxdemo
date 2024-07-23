### 1、安装
#### 1.1 下载源码
```
git clone https://github.com/ynwcel/gfxdemo 
```
#### 1.2 重命名包名（可选）
```
go run hack/hack.go --setgomod=mygoweb
```

### 2、配置
#### 2.1 项目配置
```shell
# 方式一:项目根目录下
cp public/config.demo.yaml ./config.yaml

# 方式二: 资源目录下，（会打包进可执行文件)
cp public/config.demo.yaml public/config.yaml
```
#### 2.2 gf 命令配置
```shell
cp hack/config.demo.yaml hack/config.yaml
```

### 3、编译
#### 3.1 自动设置版本及日期
```
go run hack/hack.go --build --build.version=0.0.1
```

