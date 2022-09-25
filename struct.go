package go_chrome_build

type PackageConf struct {
	Name              string      `json:"name"`
	ChromeExecPath    string      `json:"chrome_exec_path"`
	IntegratedBrowser bool        `json:"integrated_browser"`
	ChromePackPath    PlatformAll `json:"chrome_pack_path"`
	ChromeVersion     Platform    `json:"chrome_version"`
	Icons             Platform    `json:"icons"`
	WindowsArch       string      `json:"windows_arch"`
	DarwinAppleChip   bool        `json:"darwin_apple_chip"`
}

type Platform struct {
	Linux   string `json:"linux"`
	Windows string `json:"windows"`
	Darwin  string `json:"darwin"`
}

type PlatformAll struct {
	Linux     string `json:"linux"`
	Windows64 string `json:"windows64"`
	Windows   string `json:"windows"`
	Darwin    string `json:"darwin"`
	DarwinArm string `json:"darwin_arm"`
}

type packageMsg struct {
	sysStruct string
	version   string
	osName    string
}
