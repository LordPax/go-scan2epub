package service

import (
	"os"
	"scan2epub/utils"
)

func Scan2Epub(chaps []string) error {
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	for _, chap := range chaps {
		if err := downloadChap(chap); err != nil {
			log.PrintfErr("%v\n", err)
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
