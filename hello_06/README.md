命名
====
一个名字必须以一个字母或下划线开头，后面可以跟任意数量的字母、数字或下划线
驼峰式 命名，当名字由几个单词组成时优先使用大小写分隔

声明
====
Go语言主要有四种类型的声明语句：var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明

变量声明
-------
var 变量名字 类型
名字 := 表达式

指针
----
``` go
p := &x //p指针指向变量x
*p = 2 //*p表达式读取指针指向的变量的值
```

new函数
-------
``` go
p := new(int)   // p, *int 类型, 指向匿名的 int 变量
fmt.Println(*p) // "0"
*p = 2          // 设置 int 匿名变量的值为 2
fmt.Println(*p) // "2"
```

复数
----
复数类型：complex64和complex128，分别对应float32和float64

执行
====
>>> go build hello.go
>>> ./hello a bc def
>>> ./hello -s / a bc def