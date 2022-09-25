package go_chrome_build

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type downloader struct {
	io.Reader
	Total   int64
	Current int64
}

func (d *downloader) Read(p []byte) (n int, err error) {
	n, err = d.Reader.Read(p)
	d.Current += int64(n)
	fmt.Printf("\rDownloading... %.2f%%", float64(d.Current*10000/d.Total)/100)
	if d.Current == d.Total {
		fmt.Printf("\rDownload completed %.2f%%", float64(d.Current*10000/d.Total)/100)
	}
	return
}

func downloadHandle(url, filePath string, wg *sync.WaitGroup) error {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	downloaderMsg := &downloader{
		Reader: resp.Body,
		Total:  resp.ContentLength,
	}
	_, err = io.Copy(file, downloaderMsg)
	return err
}

func DownloadFile(url, filePath string) error {
	var wg sync.WaitGroup
	task := make(map[string]string)
	task[url] = filePath
	for k, v := range task {
		wg.Add(1)
		err := downloadHandle(k, v, &wg)
		if err != nil {
			return err
		}
	}
	wg.Wait()
	fmt.Println("")
	return nil
}
