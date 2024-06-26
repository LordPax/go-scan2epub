package service

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"scan2epub/config"
	"strconv"
)

type Page struct {
	Url  string
	Path string
}

func isPageExist(url string) bool {
	resp, err := http.Get(url)
	return err == nil && resp.StatusCode == 200
}

func getWorkingUrl(url, page string) (string, string) {
	for _, ext := range config.AVAILABLE_EXT {
		imgURL := fmt.Sprintf("%s/%s.%s", url, page, ext)
		if isPageExist(imgURL) {
			return imgURL, ext
		}
	}

	return "", ""
}

func formatPageName(page string) string {
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 10 {
		return "0" + page
	}

	return page
}

func getListOfPages(url, pathPage string) []Page {
	var page []Page

	for i := 0; ; i++ {
		formatPage := formatPageName(strconv.Itoa(i))
		imgURL, ext := getWorkingUrl(url, formatPage)

		if imgURL == "" {
			break
		}

		fileName := fmt.Sprintf("%s.%s", formatPage, ext)
		pathName := path.Join(pathPage, fileName)
		pageFound := Page{Url: imgURL, Path: pathName}

		page = append(page, pageFound)
	}

	return page
}

func getPageFromDir(pathPage string) []string {
	var pages []string

	files, _ := os.ReadDir(pathPage)
	for _, file := range files {
		fullPath := path.Join(pathPage, file.Name())
		pages = append(pages, fullPath)
	}

	return pages
}
