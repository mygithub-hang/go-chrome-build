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

func darwinPackDarwin() {
	fmt.Println("拼命开发中...")
}

func darwinPackWindows(architecture ...string) {
	architectureName := "amd64" // 386
	if len(architecture) != 0 {
		if architecture[0] == "amd64" {

		} else if architecture[0] == "386" {
			architectureName = "386"
		} else {
			panic("architecture error")
		}
	}
	conf := getConfig()
	gp := os.Getenv("GOPATH")
	if len(gp) == 0 {
		gp = build.Default.GOPATH
	}
	args := []string{
		"build",
		"-ldflags",
		"-H windowsgui",
		//"-o",
		//fmt.Sprintf("./output/%s.exe", conf.Name),
	}
	var cmd = exec.Command("go", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Dir = conf.RunBuildPath
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Env = append(cmd.Env,
		"CGO_ENABLED=0",
		"GOOS=windows",
		"GOARCH="+architectureName,
		"GOPATH="+gp,
		//"GOARCH=",
	)
	err := cmd.Run()
	log.Println(cmd.Args)
	if err != nil {
		msg := fmt.Sprint(err) + ": " + stderr.String()
		err = errors.New(msg)
		EchoError(err.Error())
	}
	log.Println(out.String())
}

func darwinPackLinux() {
	fmt.Println("拼命开发中...")
}
