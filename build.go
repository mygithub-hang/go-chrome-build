package go_chrome_build

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func DoBuild(sysType string, architecture ...string) {
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
	if conf.IntegratedBrowser {
		// 获取浏览器位置
		browserPath, browserName := getBrowserPath(sysType)
		// 获取打包浏览器位置
		runPath := GetWorkingDirPath()
		browserDir := runPath + "/resources/browser"
		err := createDir(browserDir)
		if err != nil {
			EchoError("create dir browser error")
			return
		}
		newBrowserPath := browserDir + "/" + browserName
		//defer os.Remove(newBrowserPath)
		err = createDir(newBrowserPath)
		if err != nil {
			EchoError("copy browser dir err: " + err.Error())
		}
		if !IsExist(newBrowserPath) {
			// 打包目录不存在文件
			sysStruct := ""
			version := ""
			sys := ""
			if browserPath == "" {
				// 浏览器不存在 下载到打包目录
				switch sysType {
				case "darwin":
					version = conf.ChromeVersion.Darwin
					sys = "darwin"
					if architectureName == "386" {
						panic("MacOs does not support 386")
					} else {
						sysStruct = "Mac"
					}
				case "linux":
					version = conf.ChromeVersion.Linux
					sys = "linux"
					if architectureName == "386" {
						panic("linux does not support 386")
					} else {
						sysStruct = "Linux_x64"
					}
				case "windows":
					version = conf.ChromeVersion.Windows
					sys = "win"
					if architectureName == "386" {
						sysStruct = "Win"
					} else {
						sysStruct = "Win_x64"
					}
				}
				DownBrowser(sysStruct, version, sys, newBrowserPath)
			} else {
				// 存在直接复制
				_, err = copyFile(newBrowserPath, browserPath)
				if err != nil {
					EchoError("copy browser err: " + err.Error())
					return
				}
			}
		}
	}
	buildDir := []string{
		"./resources/...",
	}
	err := Pack(sysType+"_build.go", "main", buildDir)
	if err != nil {
		EchoError(err.Error())
	} else {
		EchoSuccess("Build Compilation complete.")
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
		browserName = "chrome-mac.zip"
		if packConfig.ChromePackPath.Darwin != "" && IsExist(packConfig.ChromePackPath.Darwin) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Darwin, "chrome-mac.zip") {
				EchoError("filename must be chrome-mac.zip")
				os.Exit(1)
			}
			browserPath = packConfig.ChromePackPath.Darwin
		}
	case "linux":
		browserName = "chrome-linux.zip"
		if packConfig.ChromePackPath.Linux != "" && IsExist(packConfig.ChromePackPath.Linux) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Linux, "chrome-linux.zip") {
				EchoError("filename must be chrome-linux.zip")
				os.Exit(1)
			}
			browserPath = packConfig.ChromePackPath.Linux
		}
	case "windows":
		browserName = "chrome-win.zip"
		if packConfig.ChromePackPath.Windows != "" && IsExist(packConfig.ChromePackPath.Windows) {
			if !strings.HasSuffix(packConfig.ChromePackPath.Windows, "chrome-win.zip") {
				EchoError("filename must be chrome-win.zip")
				os.Exit(1)
			}
			browserPath = packConfig.ChromePackPath.Windows
		}
	}
	return browserPath, browserName
}

func DownBrowser(sysStruct, version, sys, downPath string) string {
	sysStructArr := map[string]string{
		"Linux_x64": "",
		"Mac":       "",
		"Mac_Arm":   "",
		"Win":       "",
		"Win_x64":   "",
	}
	if _, ok := sysStructArr[sysStruct]; !ok {
		panic("System Architecture is not in the list;{Linux_x64,Mac,Mac_Arm,Win,Win_x64}")
	}
	downUrl := fmt.Sprintf("https://registry.npmmirror.com/-/binary/chromium-browser-snapshots/%s/%s/chrome-%s.zip",
		sysStruct, version, sys)
	if downPath == "" {
		downPath = GetWorkingDirPath() + fmt.Sprintf("/browser/chrome-%s.zip", sys)
	}
	err := DownloadFile(downUrl, downPath+".tmp")
	if err != nil {
		_ = os.Remove(downPath + ".tmp")
		panic(err)
	} else {
		_ = os.Rename(downPath+".tmp", downPath)
	}
	return downPath
}
