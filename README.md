# MGgolang
## 这个仓库是我的 Go 语言学习项目。

<u>*空切片与nil切片：*</u>

1. []int{} => 空切片
2. var slice []int => nil切片

<u>*命令行的用户管理：*</u>

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


*<u>包（package）：</u>*

最小的代码分发单位，可作为独立单元维护。

放置函数、变量、结构体、常量。

用以区分库文件和可运行程序，隔离包内和包外的变量的访问关系。

main包，是可运行的，可编译为可执行程序。

非main包，作为库文件，用以提供功能。

<u>*单元测试：*</u>

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

<u>*基准测试：*</u>

> [!IMPORTANT]
>
> 性能测试：函数名写为BenchmarkXxx(b *testing.B)，测试代码的运行效率。
>
> ```
> // 代码的性能测试
> $ go test Xxx.go -bench . -benchmem
> ```

