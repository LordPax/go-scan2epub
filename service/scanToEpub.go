package service

import (
	"os"
)

func Scan2Epub(chaps []string) error {
	for _, chap := range chaps {
		if err := downloadChap(chap); err != nil {
			return err
		}
		if err := convertChap(chap); err != nil {
			return err
		}
	}

	return nil
}

func CheckChapExist(chap string) bool {
	url := os.Getenv("URL")
	urlChap := url + "/" + chap
	workingUrl, _ := getWorkingUrl(urlChap, "00")

	return workingUrl != ""
}
