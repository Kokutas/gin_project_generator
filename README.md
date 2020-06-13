# gin 项目生成器

# 编译
## 交叉编译
```shell script
# Windows x64(64位Windows操作系统)
#set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -o gin_project_genetator_windows_amd64.exe main.go

# Windows x86(32位Windows操作系统)
#set CGO_ENABLED=0
set GOOS=windows
set GOARCH=386
go build -o gin_project_genetator_windows_386.exe main.go

# Linux x64(64位Linux操作系统)
#set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o gin_project_genetator_linux_amd64 main.go

# Linux x86(32位Linux操作系统)
#set CGO_ENABLED=0
set GOOS=linux
set GOARCH=386
go build -o gin_project_genetator_linux_386 main.go

```
# 使用
```shell
Example:
#-----------------------------------RUN----------------------------------#
# Windws amd64 :  gin_project_genetator_windows_amd64.exe <project_name> #
# Windws 386 :      gin_project_genetator_windows_386.exe <project_name> #
# Linux amd64 :         gin_project_genetator_linux_amd64 <project_name> #
# Linux 386 :             gin_project_genetator_linux_386 <project_name> #
#------------------------------------------------------------------------#
```
## 生成目录结构
```shell
project/
├── common
├── config
│   ├── database.toml
│   ├── nosql.toml
│   └── server.toml
├── controller
├── dao
├── doc
├── dto
├── go.mod
├── init
├── main.go
├── model
├── proto
│   └── pb
├── README.md
├── router
│   └── middleware
├── service
├── static
├── template
├── test
├── util
└── vendor
```
<table><thead><th>生成项</th><th>说明</th></thead>
<tbody>
<tr><td>project</td><td>项目名</td></tr><tr>
<td>common</td><td>公共定义的常量/返回码等</td></tr>
<tr><td>config</td><td>项目配置文件</td></tr>
<tr><td>config/database.toml</td><td>数据库配置文件：mysql/oracle/sqllite3等</td></tr>
<tr><td>config/nosql.toml</td><td>缓存配置文件：redis/MongoDB/etcd等配置</td></tr>
<tr><td>config/server.toml</td><td>系统配置文件：开发ip/port，测试ip/port,模式[开发、测试、生产]设定等等</td></tr>
<tr><td>controller</td><td>控制器：验证提交的数据，将验证后的数据传递给service[MVC中的C]</td></tr>
<tr><td>dao</td><td>数据库操作：不写业务逻辑，只进行数据库操作：[DAO:data access object]数据访问对象</td></tr>
<tr><td>doc</td><td>文档：Swagger/go doc文档</td></tr> 
<tr><td>dto</td><td>系统服务交互传输：[Data Transfer Object]数据传输对象，表示层[MVC中的v]数据交互的结构体，可以是model中的一部分，比如json</td></tr> 
<tr><td>go.mod</td><td>go mod 文件</td></tr> 
<tr><td>init</td><td>系统初始化:环境检测，默认配置设置</td></tr>  
<tr><td>main.go</td><td>main 文件</td></tr> 
<tr><td>model</td><td>数据库ORM[object relationship mapping]对象关系映射[MVC中的M]，也可以写成entity</td></tr>
<tr><td>proto</td><td>远程过程调用[RPC:remote process call]的grpc的protobuf文件</td></tr>  
<tr><td>proto/pb</td><td>远程过程调用[RPC:remote process call]的grpc的*.pb.go</td></tr>  
<tr><td>README.md</td><td>自述文件</td></tr>  
<tr><td>router</td><td>路由：路由配置及路由的中间件（鉴权、日志、异常捕获）</td></tr>  
<tr><td>router/middleware</td><td>中国件</td></tr>  
<tr><td>service</td><td>业务：完成业务逻辑的开发，不进行数据库的操作</td></tr> 
<tr><td>static</td><td>如果没有做前后端分离的静态资源文件目录：js/css/images...</td></tr> 
<tr><td>template</td><td>如果没有做前后端分离的模板绑定目录：index.tpl/index.html等文件,也可以写成view[MVC中的V]</td></tr> 
<tr><td>test</td><td>测试目录：主要指httptest交互性测试等文件</td></tr> 
<tr><td>util</td><td>通用工具类</td></tr> 
<tr><td>vendor</td><td>第三方扩展包,go vendor管理</td></tr> 
</tbody>
</table>
