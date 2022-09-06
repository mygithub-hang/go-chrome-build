package go_chrome_build

import (
	"github.com/go-bindata/go-bindata/v3"
	"path/filepath"
	"strings"
)

func Pack(packPath, packName string, dir []string) error {
	cfg := bindata.NewConfig()
	cfg.Output = packPath
	cfg.Package = packName
	cfg.Input = []bindata.InputConfig{}
	cfg.HttpFileSystem = true
	cfg.Prefix = "resources/"
	for _, v := range dir {
		cfg.Input = append(cfg.Input, parseInput(v))
	}

	err := bindata.Translate(cfg)
	return err
}

func parseInput(path string) bindata.InputConfig {
	if strings.HasSuffix(path, "/...") {
		return bindata.InputConfig{
			Path:      filepath.Clean(path[:len(path)-4]),
			Recursive: true,
		}
	} else {
		return bindata.InputConfig{
			Path:      filepath.Clean(path),
			Recursive: false,
		}
	}
}
