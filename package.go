package go_chrome_build

var confTemp = `{
    "name": "go-chrome-demo",
    "chrome_exec_path": "",
    "integrated_browser": true,
    "chrome_pack_path": {
        "linux": "%s/browser/Linux_x64/chrome-linux.zip",
        "windows64": "%s/browser/Win_x64/chrome-win.zip",
        "windows": "%s/browser/Win/chrome-win.zip",
        "darwin": "%s/browser/Mac/chrome-mac.zip",
        "darwin_arm": "%s/browser/Mac_Arm/chrome-mac.zip"
    },
    "chrome_version": {
        "linux": "",
        "windows": "",
        "darwin": ""
    },
    "icons": {
        "linux": "resources/icons/icon.png",
        "windows": "resources/icons/icon.ico",
        "darwin": "resources/icons/icon.icns"
    },
    "windows_arch": "amd64",
    "darwin_apple_chip": false
}
`
