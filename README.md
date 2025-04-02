# MGgolang
## 这个仓库是我的 Go 语言学习项目。

### <u>*空切片与nil切片*</u>

1. []int{} => 空切片
2. var slice []int => nil切片

### <u>*命令行的用户管理*</u>

1. 用户信息存储

    > [!TIP]
    >
    > => 内存
    >
    > => 结构 []  map
    >
    > => 用户 ID name age tel addr
    >
    > map

2. 用户添加

3. 用户查询

4. 用户修改
   
   > [!TIP]
   >
   >  // 输入需要修改的用户ID：
   >
   >  // user[id] => 在 不在（输入的不正确）
   >
   >  // 打印用户信息，提示用户是否确认修改（y/N）
   >
   >  // y 提示用户输入修改后的内容 
   >
   >  // name，age，tel，addr
   
5. 用户删除
   
   > [!TIP]
   >
   > // 请输入需要删除的用户ID：
   >
   > // user[id] => 在 不在（输入的不正确）
   >
   > // 打印用户信息，提示用户是否确认修改（y/N）
   >
   > // y，就delete()
   
6. 设置登录验证

   > [!TIP]
   >
   > 在程序中定义password = "123abc!@#"
   >
   > 提示输入密码，如果密码输入3次都失败，提示并退出
   >
   > 如果密码正确，再进行用户管理操作

### *<u>包（package）</u>*

最小的代码分发单位，可作为独立单元维护。

放置函数、变量、结构体、常量。

用以区分库文件和可运行程序，隔离包内和包外的变量的访问关系。

main包，是可运行的，可编译为可执行程序。

非main包，作为库文件，用以提供功能。

### <u>*单元测试*</u>

> [!IMPORTANT]
>
> 1. 当要测试某个Xxx.go的功能的时候，在其路径下创建Xxx_test.go文件，引入testing库。
> 2. 功能测试：函数名写为TestXxx(t  *testing.T)，函数体按照自己的需求写入测试内容。
>
>
> ```
> // 运行测试结果
> $ go test Xxx.go
> // 测试覆盖率
> $ go test Xxx.go -coverprofile=cover.out
> // 显式展示测试覆盖情况
> $ go tool cover -html cover.out
> ```
>
> 

### <u>*基准测试*</u>

> [!IMPORTANT]
>
> 性能测试：函数名写为BenchmarkXxx(b *testing.B)，测试代码的运行效率。
>
> ```
> // 代码的性能测试
> $ go test Xxx.go -bench . -benchmem
> ```

### *<u>struct</u>*

对若干属性（变量名+类型名）的封装

```
结构体可见性：名字首字母大小写

StructName 包外可见

structName 包外不可见
```

```
属性可见性：名字首字母大小写

AttrName 包外可见

attrName 包外不可见
```

*<u>匿名结构体：</u>*直接创建结构体的变量（赋值），var StructName struct{ fileds... } {value...}。访问：varname.filedname= value

### *<u>匿名嵌套</u>*

> [!IMPORTANT]
>
> 大写结构名，大写属性名 => 嵌入 => 小写结构体 => 包外不能创建结构体对象
>
> 大写结构名，小写属性名 => 嵌入 => 小写结构体 => 包外不能创建结构体对象
>
> 大写结构名，大写属性名 => 嵌入 => 大写结构体 => 包外可创建结构体对象，属性可访问
>
> 大写结构名，小写属性名 => 嵌入 => 大写结构体 =>  包外可创建结构体对象，属性不可访问
>
> *小写结构名，大写属性名 => 嵌入 => 大写结构体 =>  包外可创建结构体对象，属性可访问（特殊）*
>
> 小写结构名，小写属性名 => 嵌入 => 大写结构体 =>  包外可创建结构体对象，属性不可访问
>
> 小写结构名，大写属性名 => 嵌入 => 小写结构体 =>  包外不能创建结构体对象
>
> 小写结构名，小写属性名 => 嵌入 => 小写结构体 =>  包外不能创建结构体对象

### <u>*方法值&方法表达式*</u>

方法值： Func := 变量名.方法名 => 调用：func( )

方法表达式： func := 结构体名.方法名 => 调用：func(结构体变量, 其他变量)。

### *<u>文件IO</u>*

```
读：
   a. 文件位置
   b. 打开文件
   c. 读取文件内容
   d. 关闭文件

文件路径：
   绝对路径：程序不管位置如何，操作的文件都不会变化（从根路径/盘符开始书写的路径）
   相对路径：与程序的位置有关（程序运行的位置）./xxxx   ../xxxx    xxxx
			程序放置的位置：(例){$HOME}/Download/go/main
			程序运行的位置：(例)「当前路径」{$HOME}/Library  「文件路径」{$HOME}/Download/go/main

```

```
写：
   a. 创建文件
   b. 写入内容
   c. 关闭文件
```

```
追加：
	 // append
   在文件的内容末尾加上新的内容
   比如：打开一个txt文件，在文本的末尾添加新的文本。
```

```
os.openfile <--> 磁盘
strings.Reader/Writer/Builder <--> 内存
bufio ：以上2者的扩展
```

### *<u>接口Interface</u>*

是自定义类型，对其他类型的行为的抽象。强调各个类型的行为，对相同行为进行统一。

### <u>*反射*</u>



### *<u>goroutine</u>*

```
并发，一个时间点只能有一个例程在跑。

格式：go func()

示例： main函数，为“主例程”。go func()启动的例程，为“工作例程”。

主程序的执行结束时间若比goroutine早，goroutine就自动结束。

可用sync.WaitGroup(计数信号量)来维护执行例程的执行状态。
也可用runtime.GoSched()使例程主动让出CPU。
还可用time.Sleep(time.second)让例程休眠，从而让出CPU。

锁：
Lock
Unlock
对资源操作的同时获取锁
获取了 => 操作 => 释放锁
没获取 => 等待

读写锁：
lock
RLock

Unlock
RUnlock
n个在写，m个在读
多个读，不需要锁
1读 / 写
RLock => RLock 不阻塞
RLock => Lock 阻塞 （要使用RUnlock）
Lock => RLock 阻塞
Lock => Lock 阻塞 （要使用Unlock）
```

### <u>*网络编程*</u>

服务器：监听某端口，设置一个ip，等待客户端连接。连接成功后向客户端发送成功确认。



### *<u>HTTP</u>*

```
Request: 
	1: 请求行 \r\n	GET/POST url HTTP/1.0
	2-n: key: value \r\n 请求头
	\r\n
	请求体
```

```
Response:
	1: 响应行 HTTP/1.0 STATUS_CODE STATUS_TEXT
	2-n: key: value 响应头
	\r\n
	响应体
```

