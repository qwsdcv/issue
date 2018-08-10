# issue


## Lib

```
go get -u github.com/golang/dep/cmd/dep
dep init
dep status
dep ensure
```


## 备忘
### OpenUi5

添加上`data-sap-ui-bindingSyntax="complex"`在html，否则不能使用formatter

SplitApp ， to/toDetail/toMaster 必须用 this.createId(id) 创建id作参数，否则不work

### 交叉编译

在windows上

```
set GOOS=linux
set GOARCH=amd64
go build
```