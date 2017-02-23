# 简介

golang的基础数据类型相互转换包。

包括ToXX方法及MustXX两类方法，To方法会在失败的时候返回error，而Must方法会在失败的时候直接panic。

# 安装

```
go get -u -v go.szyhf.org/di-convert
```

# 例子

```golang
import convert"go.szyhf.org/di-convert"
import "fmt"

var numStr := "123"
num,err := convert.GetInt(numStr)
fmt.Println(num,err)

num2 := convert.MustInt(numStr)
fmt.Println(num)
```