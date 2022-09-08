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

}

func windowsPackWindows() {
	conf := getConfig()
	gp := os.Getenv("GOPATH")
	if len(gp) == 0 {
		gp = build.Default.GOPATH
	}
	args := []string{
		"build",
		"-ldflags",
		"-H windowsgui",
	}
	var cmd = exec.Command("go", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Env = os.Environ()
	cmd.Dir = conf.RunBuildPath
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

}
