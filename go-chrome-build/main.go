package main

import (
	"fmt"
	go_chrome_build "github.com/mygithub-hang/go-chrome-build"
	"os"
	"strings"
)

// 控制台颜色文字变量
var (
//cyan  = string([]byte{27, 91, 51, 54, 109})
//reset = string([]byte{27, 91, 48, 109})
)

// go get github.com/mygithub-hang/go-chrome-build
// go install -a -v github.com/mygithub-hang/go-chrome-build/...
// init 初始化
func init() {
	//flags := true
	//if !IsExist("./config/config.toml") {
	//	EchoError("当前目录下未找到：config/config.toml 文件")
	//	flags = false
	//}
	//if !flags {
	//	EchoError("请在完整的gf框架根目录执行命令")
	//	EchoError("5秒后自动退出。。。")
	//	time.Sleep(5 * time.Second)
	//	exit()
	//}
}

func main() {
	param := os.Args
	buildCmd := ""
	if len(param) > 1 {
		buildCmd = param[1]
	}
	Run(buildCmd)
}

func Run(cmd string) {
	if cmd != "b" && cmd != "w" && cmd != "l" && cmd != "d" {
		help()
		terminal()
	} else {
		fmt.Println(go_chrome_build.GetCurrentPath())
		fmt.Println(go_chrome_build.GetExcPath())
		fmt.Println(cmd)
	}
}

func help() {
	helpText := `
---------------HELP---------------
 b 		编译当前系统可执行文件
 w 		编译win X86可执行文件
 l 		编译linux X86可执行文件
 d 		编译MacOs X86可执行文件
 h,help 	帮助
 exit		退出
---------------HELP---------------`
	fmt.Println(go_chrome_build.Cyan, helpText, go_chrome_build.Reset)
}

func terminal() {
	cmdStr, err := go_chrome_build.AskForConfirmation(""+
		"go-chrome-build>", "")
	if err != nil {
		go_chrome_build.EchoError(err.Error())
	}
	if cmdStr == "exit" {
		exit()
	}
	newCmdStr := strings.Trim(cmdStr, " ")
	Run(newCmdStr)
}

// exit 退出控制台
func exit() {
	os.Exit(0)
}
