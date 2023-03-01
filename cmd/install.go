package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"strings"
)

// 接收输入参数：--arg
var hosts string
var hostPassword string
var connectPassword string
var maxMemory int64

func init() {
	//绑定输入参数：--arg
	InstallCmd.Flags().StringVarP(&hosts, "hosts", "i", "127.0.0.1", "输入IP,多个采用逗号分割")
	InstallCmd.Flags().StringVarP(&hostPassword, "hostPassword", "P", "1", "主机密码")
	InstallCmd.Flags().StringVarP(&connectPassword, "connectPassword", "p", "jjy522", "连接密码")
	InstallCmd.Flags().Int64VarP(&maxMemory, "maxMemory", "m", 1073741824, "内存最大值")
}

var InstallCmd = &cobra.Command{
	Use:     "install",
	Version: "1.0.0",
	Short:   "Install command",
	Long:    "该子命令用于启动ansible部署",
	Run: func(cmd *cobra.Command, args []string) {
		//参数处理
		replaceHosts := strings.Replace(hosts, ",", " ", -1)
		command := fmt.Sprintf("sh install.sh %s %s %s %d", replaceHosts, hostPassword, connectPassword, maxMemory)
		//1、命令执行
		//err := execCommand("ping -c 5 baidu1.com")
		//2、命令执行
		err := execCommandLine(command)
		if err != nil {
			exitError := err.(*exec.ExitError)
			os.Exit(exitError.ExitCode())
		}
	},
}

// 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func execCommand(command string) error {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", command)
	//显示运行的命令
	fmt.Println(cmd.Args)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	fmt.Println(out.String())
	return err
}

// 阻塞式的执行外部shell命令的函数,逐行实时进行处理的
func execCommandLine(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}
