# Deploy-CLI
>基于Cobra实现的CLI应用程序Demo

# 构建
- `go mod tidy`把项目所依靠的包增加到go.mod文件中
- `go run deploy-cli.go`调试执行
- `go build deploy-cli.go `编译成`deploy-cli`

# 示例
```shell
➜  deploy-cli go run deploy-cli.go
Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install command

Flags:
  -h, --help      help for this command
  -v, --version   version for this command

Use " [command] --help" for more information about a command.
```
```shell
➜  deploy-cli go run deploy-cli.go install -h                              
该子命令用于启动ansible部署

Usage:
   install [flags]

Flags:
  -h, --help              help for install
  -i, --ip string         输入IP,多个采用逗号分割 (default "127.0.0.1")
  -p, --password string   输入密码 (default "123456")
  -n, --port int32        输入端口 (default 8080)
  -v, --version           version for install

```

# 注意：
```go
//比如命令：./deploy-cli sub arg1 arg2 --arg3=v1 --arg4=v4 
var subCmd = &cobra.Command{
    Use:   "sub",
    Short: "My subcommand",
    Run: func(cmd *cobra.Command, args []string) {
      //这里的获取的args []string是指arg1 arg2，并不是arg3 arg4
      fmt.Printf("Inside subCmd Run with args: %v\n", args)
    },
  }
}        
```