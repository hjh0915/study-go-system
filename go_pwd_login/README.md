Go 语言使用 Bcrypt 实现加密与验证
==============================
``` go
mkdir pwd_demo
go mod init pwd_demo
```

引入
----
``` go
import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)
```

加密
----
bcrypt.GenerateFromPassword()同bcrypt中的hashpw(pwd, salt)

验证
----
bcrypt.CompareHashAndPassword()同bcrypt中的checkpw(pwd1, pwd2)

参考链接：https://blog.csdn.net/cnwyt/article/details/119460737

使用go test测试
==============
1、测试文件以 _test.go 结尾 
2、测试用例名称一般命名为 Test 加上待测试的方法名。 
3、测试用的参数有且只有一个，在这里是 t *testing.T。 
4、调用方法进行测试 

执行
----
执行测试代码(只返回结果): 
$ go test utils/pwd_test.go utils/pwd.go 
 
查看执行过程和测试代码里的日志: 
$ go test -v utils/pwd_test.go utils/pwd.go