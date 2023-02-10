package main

import (
	"deploy-cli/cmd"
	"github.com/spf13/cobra"
	"log"
)

// root命令
var rootCmd = &cobra.Command{Version: "1.0.0"}

func init() {
	//添加子命令
	rootCmd.AddCommand(cmd.InstallCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
