package go_chrome_build

import (
	"os"
	"runtime"
	"strings"
)

func DoBuild(sysType string) {
	browserPath, browserName := getBrowserPath(sysType)
	runPath := GetCurrentPath()
	browserDir := runPath + "/browser"
	err := createDir(browserDir)
	if err != nil {
		EchoError("create dir browser error")
		return
	}
	newBrowserPath := browserDir + "/" + browserName
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
		EchoSuccess("Compilation complete.")
	}
}

func PackWindows() {
	DoBuild("windows")
	err := BuildSysO()
	if err != nil {
		EchoError(err.Error())
		os.Exit(0)
		return
	}
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
		if packConfig.ChromePackPath.Darwin != "" && IsExist(packConfig.ChromePackPath.Darwin) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Darwin, "chrome-mac.zip") {
				EchoError("filename must be chrome-mac.zip")
				os.Exit(1)
			}
			browserName = "chrome-mac.zip"
			browserPath = packConfig.ChromePackPath.Darwin
		}
	case "linux":
		if packConfig.ChromePackPath.Linux != "" && IsExist(packConfig.ChromePackPath.Linux) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Linux, "chrome-linux.zip") {
				EchoError("filename must be chrome-linux.zip")
				os.Exit(1)
			}
			browserName = "chrome-linux.zip"
			browserPath = packConfig.ChromePackPath.Linux
		}
	case "windows":
		if packConfig.ChromePackPath.Windows != "" && IsExist(packConfig.ChromePackPath.Windows) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Windows, "chrome-win.zip") {
				EchoError("filename must be chrome-win.zip")
				os.Exit(1)
			}
			browserName = "chrome-win.zip"
			browserPath = packConfig.ChromePackPath.Windows
		}
	}
	return browserPath, browserName
}
