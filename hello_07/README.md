数组
====
数组是一个由固定长度的特定类型元素组成的序列
和数组对应的类型是Slice（切片），它是可以增长和收缩的动态序列
``` go
var q [3]int = [3]int{1, 2, 3}
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // "0"

q := [...]int{1, 2, 3}
q := [3]int{1, 2, 3}
```

指针
====
指针来传递数组参数是高效的

切片
====
一个slice类型一般写作[]T，由指针、长度和容量组成。多个slice之间可以共享底层的数据
slice唯一合法的比较操作是和nil比较
测试一个slice是否是空的，使用len(s) == 0来判断

map
====
创建方法：
--------
1、使用make函数
``` go
ages := make(map[string]int)
```
2、直接使用map创建
``` go 
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```
创建空的map: map[string]int{}, 在向map存数据前必须先创建map。
*禁止对map元素取址: map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。*

结构体
======
*可嵌套*
对成员取地址，然后通过指针访问
--------------------------
``` go
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee

position := &dilbert.Position
*position = "Senior " + *position
```

点操作符和指向结构体的指针
----------------------
``` go 
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
```
*S类型的结构体可以包含S指针类型的成员*

指定值
------
``` go
type Point struct{ X, Y int }
p := Point{1, 2}
```

较大的结构体通常会用指针的方式传入和返回
----------------------------------
``` go
func Bonus(e *Employee, percent int) int {
    return e.Salary * percent / 100
}
```

JSON
=====
调用json.Marshal函数实现结构体slice转为JSON的过程