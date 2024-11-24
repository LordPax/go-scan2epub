package service

import (
	"scan2epub/config"
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
	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	url := config.CONFIG_INI.Section(defaultSource).Key("url").String()
	startAt, err := config.CONFIG_INI.Section(defaultSource).Key("start_at").Int()
	if err != nil {
		return false
	}

	formatPage := formatPageName(strconv.Itoa(startAt))
	workingUrl, _ := getWorkingUrl(url, chap, formatPage)

	return workingUrl != ""
}

func CronDownloadChap() error {
	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	cronStr := config.CONFIG_INI.Section(defaultSource).Key("cron").String()
	currentChap, err := config.CONFIG_INI.Section(defaultSource).Key("cron_chap").Int()
	if err != nil {
		return err
	}

	c, err := cron.NewScheduler()
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
		log.PrintfErr(l.Get("chapter-not-found"), *currentChap)
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
