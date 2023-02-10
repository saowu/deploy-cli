package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 接收输入参数：--arg
var port int32
var password string
var ip string

var InstallCmd = &cobra.Command{
	Use:     "install",
	Version: "1.0.0",
	Short:   "Install command",
	Long:    "该子命令用于启动ansible部署",
	Args: func(cmd *cobra.Command, args []string) error {
		//Flags参数验证
		fmt.Printf("Inside subCmd Args with args: %v\n", args)
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//命令执行前
		fmt.Printf("Inside subCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//命令执行
		fmt.Printf("ip is %s, port is %d, password is %s \n", ip, port, password)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//命令执行后
		fmt.Printf("Inside subCmd PostRun with args: %v\n", args)
	},
}

func init() {
	//绑定输入参数：--arg
	InstallCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "输入IP,多个采用逗号分割")
	InstallCmd.Flags().Int32VarP(&port, "port", "n", 8080, "输入端口")
	InstallCmd.Flags().StringVarP(&password, "password", "p", "123456", "输入密码")
}
