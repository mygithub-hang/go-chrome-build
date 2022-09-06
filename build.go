package go_chrome_build

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func DoBuild(sysType string) {
	browserPath, browserName := getBrowserPath(sysType)
	err := createDir("./browser")
	if err != nil {
		EchoError("create dir browser error")
		return
	}
	newBrowserPath := "./browser/" + browserName
	defer os.Remove(newBrowserPath)
	_, err = copyFile(newBrowserPath, browserPath)
	if err != nil {
		EchoError("copy browser err: " + err.Error())
		return
	}
	err = Pack(sysType+"_build.go", "main", []string{
		"./resources/...",
		"./browser/...",
	})
	if err != nil {
		EchoError(err.Error())
	} else {
		fmt.Println("Compilation complete.")
	}
}

func PackWindows() {
	DoBuild("windows")
	sysType := runtime.GOOS
	switch sysType {
	case "darwin":
		// macos
		darwinPackWindows()
	case "linux":
		linuxPackWindows()
	case "windows":
		windowsPackWindows()
	}
}

func PackLinux() {
	DoBuild("linux")
	sysType := runtime.GOOS
	switch sysType {
	case "darwin":
		// macos
		darwinPackLinux()
	case "linux":
		linuxPackLinux()
	case "windows":
		windowsPackLinux()
	}
}

func PackMacOs() {
	DoBuild("darwin")
	sysType := runtime.GOOS
	switch sysType {
	case "darwin":
		// macos
		darwinPackDarwin()
	case "linux":
		linuxPackDarwin()
	case "windows":
		windowsPackDarwin()
	}
}

func PackNowSys() {
	sysType := runtime.GOOS
	switch sysType {
	case "darwin":
		// macos
		PackMacOs()
	case "linux":
		PackLinux()
	case "windows":
		PackWindows()
	}
}

func getBrowserPath(sysType string) (string, string) {
	packConfig := getConfig()
	browserPath := ""
	browserName := ""
	switch sysType {
	case "darwin":
		if packConfig.ChromePackPath.Darwin != "" && IsFile(packConfig.ChromePackPath.Darwin) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Darwin, "chrome-mac.zip") {
				EchoError("filename must be chrome-mac.zip")
				os.Exit(1)
			}
			browserName = "chrome-mac.zip"
			browserPath = packConfig.ChromePackPath.Darwin
		}
	case "linux":
		if packConfig.ChromePackPath.Linux != "" && IsFile(packConfig.ChromePackPath.Linux) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Linux, "chrome-linux.zip") {
				EchoError("filename must be chrome-linux.zip")
				os.Exit(1)
			}
			browserName = "chrome-linux.zip"
			browserPath = packConfig.ChromePackPath.Linux
		}
	case "windows":
		if packConfig.ChromePackPath.Windows != "" && IsFile(packConfig.ChromePackPath.Windows) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Linux, "chrome-win.zip") {
				EchoError("filename must be chrome-win.zip")
				os.Exit(1)
			}
			browserName = "chrome-win.zip"
			browserPath = packConfig.ChromePackPath.Windows
		}
	}
	return browserPath, browserName
}
