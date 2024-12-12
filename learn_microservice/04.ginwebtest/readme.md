踩坑：从0到1搭建 go-micro,protoc-gen-micro环境,consul作为服务发现
参考文章: https://blog.csdn.net/xuehu96/article/details/123610827
         https://blog.csdn.net/xuehu96/article/details/123617124

安装go-micro：go install go-micro.dev/v4/cmd/micro@v4.6.0
安装protoc-gen-micro：go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

到路径下去创建服务：
micro new service helloworld
cd helloworld

修改proto/xxx.proto文件
默认的proto带了三个服务，分别是call, stream, pingpong，我们用不到这么多，先把rpc里关于stream和pingpong的删掉，再把message关于stream和pingpong的删掉，只留下和call相关的

修改handler/**.go文件
只留下和Call相关的

修改 go.mod第三行，改成go 1.23

make proto
123
make tidy

在main.go中添加代码（为方便查看，添加的代码后面都用"// add做标记"）
"github.com/asim/go-micro/plugins/registry/consul/v4"包报错，最好是在vscode中手动点击添加，用go get 会出错（不知道为什么）


在新终端中启动consul：
consul agent -dev 
然后可以在[localhost](http://localhost:8500/ui/dc1/services)中打开链接

然后运行go run main.go后，在网页链接上可以看到这个服务

到此，micro服务端就弄好了

-----------------------------------------------------------
用gin搭建一个web端，并和刚才我们搭的这个服务端连通
目录树：
	hellomicro  (服务端)
	ginwebtest	(客户端)

cd ginwebtest 
go mod init
go get -u -v github.com/gin-gonic/gin   下载gin框架库
创建main.go 写代码

main.go中，最上面注释的是gin框架简单实例代码，下面才是对接micro服务端的代码

将hellomicro项目中的proto文件夹拷过来，然后.pb.micro.go中3个依赖会报错，此时，修改go.mod 文件，require添加一行go-micro.dev/v4 v4.2.1，然后go mod tidy


导入"github.com/asim/go-micro/plugins/registry/consul/v4"报错时，go mod tidy
