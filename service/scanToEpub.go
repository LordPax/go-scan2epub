package service

import (
	"fmt"
	"os"
)

func Scan2Epub(chaps []string) error {
	for _, chap := range chaps {
		if err := downloadChap(chap); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
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
