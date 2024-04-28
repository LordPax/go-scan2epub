package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"scan2epub/utils"
)

func downloadChap(chap string) error {
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	log.Printf("Downloading chapter %s\n", chap)

	channel := make(chan error)
	url := os.Getenv("URL")
	tempDir := os.Getenv("TMP_DIR")
	urlChap := url + "/" + chap
	pathChap := path.Join(tempDir, chap)

	if !utils.FileExist(pathChap) {
		if err := os.MkdirAll(pathChap, 0755); err != nil {
			return err
		}
	}

	pages := getListOfPages(urlChap, pathChap)
	log.Printf("%d pages found\n", len(pages))

	if len(pages) == 0 {
		return fmt.Errorf("No pages found for chapter %s", chap)
	}

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
	log, err := utils.GetLog()
	if err != nil {
		ch <- err
		return
	}

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

	log.Printf("Downloaded page from %s\n", url)
	ch <- nil
}
