package go_chrome_build

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
)

func windowsPackDarwin() {
	fmt.Println("拼命开发中...")
}

func windowsPackWindows(architecture ...string) {
	architectureName := "amd64" // 386
	if len(architecture) != 0 {
		if architecture[0] == "amd64" {

		} else if architecture[0] == "386" {
			architectureName = "386"
		} else {
			panic("architecture error")
		}
	}
	//conf := getConfig()
	gp := os.Getenv("GOPATH")
	if len(gp) == 0 {
		gp = build.Default.GOPATH
	}
	args := []string{
		"build",
		"-ldflags",
		"-H windowsgui",
	}
	fmt.Println("build Windows " + architectureName)
	var cmd = exec.Command("go", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env,
		"CGO_ENABLED=0",
		"GOOS=windows",
		"GOARCH="+architectureName,
		"GOPATH="+gp,
		//"GOARCH=",
	)
	cmd.Dir = GetWorkingDirPath()
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	log.Println(cmd.Args)
	if err != nil {
		msg := fmt.Sprint(err) + ": " + stderr.String()
		err = errors.New(msg)
		EchoError(err.Error())
	}
	log.Println(out.String())
}

func windowsPackLinux() {
	fmt.Println("拼命开发中...")
}
