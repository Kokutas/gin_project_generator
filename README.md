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