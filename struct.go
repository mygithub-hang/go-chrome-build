package go_chrome_build

type PackageConf struct {
	Name           string   `json:"name"`
	ChromeExecPath string   `json:"chrome_exec_path"`
	ChromePackPath Platform `json:"chrome_pack_path"`
	ChromeVersion  int      `json:"chrome_version"`
	BuildCachePath string   `json:"build_cache_path"`
	Icons          Platform `json:"icons"`
	RunBuildPath   string   `json:"run_build_path"`
}

type Platform struct {
	Linux   string `json:"linux"`
	Windows string `json:"windows"`
	Darwin  string `json:"darwin"`
}
