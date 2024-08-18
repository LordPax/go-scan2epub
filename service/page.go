package service

import (
	"net/http"
	"os"
	"path"
	"scan2epub/config"
	"strconv"
	"strings"
)

type Page struct {
	Url  string
	Path string
}

func isPageExist(url string) bool {
	resp, err := http.Get(url)
	return err == nil && resp.StatusCode == 200
}

func getWorkingUrl(url, chap, page string) (string, string) {
	for _, ext := range config.AVAILABLE_EXT {
		imgURL := replaceValue(url, map[string]string{
			"{chap}": chap,
			"{page}": page,
			"{ext}":  ext,
		})
		if isPageExist(imgURL) {
			return imgURL, ext
		}
	}

	return "", ""
}

func replaceValue(url string, data map[string]string) string {
	newUrl := url

	for k, v := range data {
		newUrl = strings.Replace(newUrl, k, v, -1)
	}

	return newUrl
}

func formatPageName(page string) string {
	pageInt, _ := strconv.Atoi(page)

	if pageInt < 10 {
		return "0" + page
	}

	return page
}

func getListOfPages(url, chap, tmpPage string) []Page {
	var page []Page

	for i := 0; ; i++ {
		formatPage := formatPageName(strconv.Itoa(i))
		imgURL, ext := getWorkingUrl(url, chap, formatPage)

		if imgURL == "" {
			break
		}

		// fileName := fmt.Sprintf("%s.%s", formatPage, ext)
		pathName := path.Join(tmpPage, formatPage+"."+ext)
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
