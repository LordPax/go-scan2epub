package service

import (
	"os"
	"scan2epub/lang"
	"scan2epub/utils"
	"strconv"

	cron "github.com/go-co-op/gocron/v2"
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
	workingUrl, _ := getWorkingUrl(url, chap, "00")

	return workingUrl != ""
}

func CronDownloadChap(cronStr, chap string) error {
	c, err := cron.NewScheduler()
	if err != nil {
		return err
	}

	currentChap, err := strconv.Atoi(chap)
	if err != nil {
		return err
	}

	chanErr := make(chan error)

	_, err = c.NewJob(
		cron.CronJob(cronStr, false),
		cron.NewTask(func() { cronFunc(&currentChap, chanErr) }),
	)
	if err != nil {
		return err
	}

	c.Start()

	if err := <-chanErr; err != nil {
		return err
	}

	return nil
}

func cronFunc(currentChap *int, ch chan<- error) {
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		ch <- err
		return
	}

	log.Printf(l.Get("current-chapter"), *currentChap)
	if !CheckChapExist(strconv.Itoa(*currentChap)) {
		log.PrintfErr(l.Get("chapter-not-found"), currentChap)
		return
	}

	if err := downloadChap(strconv.Itoa(*currentChap)); err != nil {
		ch <- err
		return
	}

	if err := convertChap(strconv.Itoa(*currentChap)); err != nil {
		ch <- err
		return
	}

	*currentChap++
	log.Printf(l.Get("next-chapter"), *currentChap)
}
