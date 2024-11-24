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

func replaceValue(s string, data map[string]string) string {
	newStr := s

	for k, v := range data {
		newStr = strings.Replace(newStr, k, v, -1)
	}

	return newStr
}

func formatPageName(page string) string {
	pageInt, _ := strconv.Atoi(page)
	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	format, _ := config.CONFIG_INI.Section(defaultSource).Key("format").Bool()

	if format && pageInt < 10 {
		return "0" + page
	}

	return page
}

func getListOfPages(url, chap, tmpPage string) []Page {
	var page []Page

	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	startAt, _ := config.CONFIG_INI.Section(defaultSource).Key("start_at").Int()

	for i := startAt; ; i++ {
		formatPage := formatPageName(strconv.Itoa(i))
		imgURL, ext := getWorkingUrl(url, chap, formatPage)

		if imgURL == "" {
			break
		}

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
