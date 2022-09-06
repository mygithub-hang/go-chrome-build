package main

import (
	"fmt"
	go_chrome_build "github.com/mygithub-hang/go-chrome-build"
	"os"
	"runtime"
	"strings"
	"time"
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
	flags := true
	runPath := go_chrome_build.GetCurrentPath()
	if !go_chrome_build.IsFile(runPath + "/package.json") {
		go_chrome_build.EchoError("当前目录下未找到：package.json 文件")
		flags = false
	}
	if !go_chrome_build.IsDir(runPath + "/resources") {
		go_chrome_build.EchoError("当前目录下未找到：resources 目录")
		flags = false
	}
	if !flags {
		go_chrome_build.EchoError("请补全目录再执行命令")
		go_chrome_build.EchoError("5秒后自动退出。。。")
		time.Sleep(5 * time.Second)
		exit()
	}
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
	if cmd != "b" && cmd != "w" && cmd != "l" && cmd != "d" && cmd != "t" {
		help()
		terminal()
	} else {
		runPack(cmd)
	}
}

func help() {
	helpText := `
---------------HELP---------------
 t 		编译执行文件不打包
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

func runPack(cmd string) {
	switch cmd {
	case "t":
		// 编译不打包
		go_chrome_build.DoBuild(runtime.GOOS)
	case "b":
		// 打包当前系统
		go_chrome_build.PackNowSys()
	case "w":
		// 打包win
		go_chrome_build.PackWindows()
	case "l":
		// 打包linux
		go_chrome_build.PackLinux()
	case "d":
		// 打包macos
		go_chrome_build.PackMacOs()
	}
}

// exit 退出控制台
func exit() {
	os.Exit(0)
}
