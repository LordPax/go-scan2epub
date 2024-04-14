package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"scan2epub/utils"
)

func downloadChap(chap string) error {
	fmt.Printf("Downloading chapter %s\n", chap)

	channel := make(chan error)
	url := os.Getenv("URL")
	tempDir := os.Getenv("TMP_DIR")
	urlChap := url + "/" + chap
	pathChap := tempDir + "/" + chap

	if !utils.FileExist(pathChap) {
		if err := os.MkdirAll(pathChap, 0755); err != nil {
			return err
		}
	}

	pages := getListOfPages(urlChap, pathChap)
	fmt.Printf("%d pages found\n", len(pages))

	for _, page := range pages {
		go downloadPage(page.Url, page.Path, channel)
	}

	for range pages {
		if err := <-channel; err != nil {
			return err
		}
	}

	return nil
}

func downloadPage(url, path string, ch chan<- error) {
	out, err := os.Create(path)
	if err != nil {
		ch <- err
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		ch <- err
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		ch <- err
		return
	}

	fmt.Printf("Downloaded page from %s\n", url)
	ch <- nil
}
