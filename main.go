package main

import (
	"fmt"
	"gin_project_generator/utils"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

/**
 *  @Author: kokutas
 *  @Date: 2020/6/13 11:40
 *  @Description: gin 的项目生成器
**/
var projectName string

func init() {
	windowsamd64 := "gin_project_genetator_windows_amd64.exe <project_name>"
	windows386 := "gin_project_genetator_windows_386.exe <project_name>"
	linuxamd64 := "gin_project_genetator_linux_amd64 <project_name>"
	linux386 := "gin_project_genetator_linux_386 <project_name>"
	log.Printf(""+
		"\r\nExample:"+
		"\r\n#-----------------------------------RUN----------------------------------#"+
		"\r\n# %7s : %55s #"+
		"\r\n# %7s : %57s #"+
		"\r\n# %7s : %56s #"+
		"\r\n# %7s : %58s #"+
		"\r\n#------------------------------------------------------------------------#"+
		"\r\n\r\n",
		"Windws amd64", windowsamd64, "Windws 386", windows386,"Linux amd64", linuxamd64,"Linux 386", linux386)

}

func main() {
	for {
		if len(os.Args) > 1 {
			projectName = os.Args[1]
			break
		} else {
			// 如果没有输入项目名称
			log.Println("您还没有指定项目名称,请参照Example的方式重新执行!")
			return
		}
	}
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 创建项目目录
	projectDir = path.Join(projectDir, projectName)
	err = utils.CreateDir(projectDir)
	if err != nil {
		log.Fatal(err)
	}

	err = utils.CreateDir(
		// 公共定义的常量/返回码等
		path.Join(projectDir, "common"),
		// 配置：项目配置文件
		path.Join(projectDir, "config"),
		// 控制器：验证提交的数据，将验证后的数据传递给service[MVC中的C]
		path.Join(projectDir, "controller"),
		// 文档：Swagger/go doc文档
		path.Join(projectDir, "doc"),
		// 数据库操作：不写业务逻辑，只进行数据库操作：[DAO:data access object]数据访问对象
		path.Join(projectDir, "dao"),
		// 系统服务交互传输：[Data Transfer Object]数据传输对象，表示层[MVC中的v]数据交互的结构体，可以使model中的一部分，比如json
		path.Join(projectDir, "dto"),
		// 系统初始化:环境检测，默认配置设置
		path.Join(projectDir, "init"),
		// 数据库ORM[object relationship mapping]对象关系映射[MVC中的M]，也可以写成entity
		path.Join(projectDir, "model"),
		// 远程过程调用[RPC:remote process call]的grpc的protobuf文件和*.pb.go
		path.Join(projectDir, "proto"),
		// *.pb.go文件
		path.Join(projectDir, "proto/pb"),
		// 路由：路由配置及路由的中间件（鉴权、日志、异常捕获）
		path.Join(projectDir, "router"),
		// 路由中间件
		path.Join(projectDir, "router/middleware"),
		// 业务：完成业务逻辑的开发，不进行数据库的操作
		path.Join(projectDir, "service"),
		// 如果没有做前后端分离的静态资源文件目录：js/css/images...
		path.Join(projectDir, "static"),
		// 如果没有做前后端分离的模板绑定目录：index.tpl/index.html等文件,也可以写成view[MVC中的V]
		path.Join(projectDir, "template"),
		// 测试目录：主要指httptest交互性测试等文件
		path.Join(projectDir, "test"),
		// 通用工具类
		path.Join(projectDir, "util"),
		// 第三方扩展包,go vendor管理
		path.Join(projectDir, "vendor"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 创建main.go
	content := "package main\r\n\r\nfunc main() {\r\n\r\n}\r\n"
	err = utils.CreateFile(projectDir, "main.go", content)
	if err != nil {
		log.Fatal(err)
	}
	// go.mod
	version := strings.ToLower(runtime.Version())
	version = strings.ReplaceAll(version, "go", "")
	if strings.Count(version, ".") > 1 {
		//fmt.Println(strings.LastIndex(version,"."))// 获取最后一次出现的.的位置(下标)
		version = version[:strings.LastIndex(version, ".")]
	}
	content = fmt.Sprintf(
		"module %s\r\n\r\n"+
			"go %s\r\n",
		projectName, version)
	err = utils.CreateFile(projectDir, "go.mod", content)
	if err != nil {
		log.Fatal(err)
	}
	// 创建README.md的Markdown文档
	content = "# 说明\r\n" +
		"## 目录结构\r\n" +
		"# 说明\r\n## 目录结构\r\n```shell\r\nproject/\r\n├── common\r\n├── config\r\n│   ├── database.toml\r\n│   ├── nosql.toml\r\n│   └── server.toml\r\n├── controller\r\n├── dao\r\n├── doc\r\n├── dto\r\n├── go.mod\r\n├── init\r\n├── main.go\r\n├── model\r\n├── proto\r\n│   └── pb\r\n├── README.md\r\n├── router\r\n│   └── middleware\r\n├── service\r\n├── static\r\n├── template\r\n├── test\r\n├── util\r\n└── vendor\r\n```\r\n" +
		"<table><thead><th>生成项</th><th>说明</th></thead>\n<tbody>\n<tr><td>project</td><td>项目名</td></tr><tr>\n<td>common</td><td>公共定义的常量/返回码等</td></tr>\n<tr><td>config</td><td>项目配置文件</td></tr>\n<tr><td>config/database.toml</td><td>数据库配置文件：mysql/oracle/sqllite3等</td></tr>\n<tr><td>config/nosql.toml</td><td>缓存配置文件：redis/MongoDB/etcd等配置</td></tr>\n<tr><td>config/server.toml</td><td>系统配置文件：开发ip/port，测试ip/port,模式[开发、测试、生产]设定等等</td></tr>\n<tr><td>controller</td><td>控制器：验证提交的数据，将验证后的数据传递给service[MVC中的C]</td></tr>\n<tr><td>dao</td><td>数据库操作：不写业务逻辑，只进行数据库操作：[DAO:data access object]数据访问对象</td></tr>\n<tr><td>doc</td><td>文档：Swagger/go doc文档</td></tr> \n<tr><td>dto</td><td>系统服务交互传输：[Data Transfer Object]数据传输对象，表示层[MVC中的v]数据交互的结构体，可以是model中的一部分，比如json</td></tr> \n<tr><td>go.mod</td><td>go mod 文件</td></tr> \n<tr><td>init</td><td>系统初始化:环境检测，默认配置设置</td></tr>  \n<tr><td>main.go</td><td>main 文件</td></tr> \n<tr><td>model</td><td>数据库ORM[object relationship mapping]对象关系映射[MVC中的M]，也可以写成entity</td></tr>\n<tr><td>proto</td><td>远程过程调用[RPC:remote process call]的grpc的protobuf文件</td></tr>  \n<tr><td>proto/pb</td><td>远程过程调用[RPC:remote process call]的grpc的*.pb.go</td></tr>  \n<tr><td>README.md</td><td>自述文件</td></tr>  \n<tr><td>router</td><td>路由：路由配置及路由的中间件（鉴权、日志、异常捕获）</td></tr>  \n<tr><td>router/middleware</td><td>中国件</td></tr>  \n<tr><td>service</td><td>业务：完成业务逻辑的开发，不进行数据库的操作</td></tr> \n<tr><td>static</td><td>如果没有做前后端分离的静态资源文件目录：js/css/images...</td></tr> \n<tr><td>template</td><td>如果没有做前后端分离的模板绑定目录：index.tpl/index.html等文件,也可以写成view[MVC中的V]</td></tr> \n<tr><td>test</td><td>测试目录：主要指httptest交互性测试等文件</td></tr> \n<tr><td>util</td><td>通用工具类</td></tr> \n<tr><td>vendor</td><td>第三方扩展包,go vendor管理</td></tr> \n</tbody>\n</table>\r\n"
	err = utils.CreateFile(projectDir, "README.md", content)
	if err != nil {
		log.Fatal(err)
	}

	// 创建配置文件
	// 1.数据库配置文件：mysql/oracle/sqllite3/...
	err = utils.CreateFile(path.Join(projectDir, "config"), "database.toml", "# mysql/oracle/sqllite3等等")
	if err != nil {
		log.Fatal(err)
	}
	// 2.系统配置文件：开发ip/port，测试ip/port等等
	err = utils.CreateFile(path.Join(projectDir, "config"), "server.toml", "# 开发ip/port，测试ip/port等等系统相关配置")
	if err != nil {
		log.Fatal(err)
	}
	// 3.缓存配置文件：redis/MongoDB/etcd等配置
	err = utils.CreateFile(path.Join(projectDir, "config"), "nosql.toml", "# redis/MongoDB/etcd等配置")
	if err != nil {
		log.Fatal(err)
	}

}
