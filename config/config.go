package config

import (
	"fmt"
	"os"
	"scan2epub/utils"
)

var home, _ = os.UserHomeDir()

var (
	NAME           = "scan2epub"
	VERSION        = "0.0.1"
	USAGE          = "CLI tool to convert scan to epub"
	CONFIG_DIR     = home + "/.config/scan2epub"
	CONFIG_FILE    = CONFIG_DIR + "/config"
	AVAILABLE_EXT  = []string{"jpg", "jpeg", "png", "webp"}
	CONFIG_EXEMPLE = `URL=https://lelscans.net/mangas/one-piece
EPUB_DIR=` + home + "/scan2epub"
)

func InitConfig() error {
	tmpName, err := os.MkdirTemp("", "scan2epub")
	if err != nil {
		return err
	}

	os.Setenv("TMP_DIR", tmpName)

	if !utils.FileExist(CONFIG_DIR) {
		if err := os.MkdirAll(CONFIG_DIR, 0755); err != nil {
			return err
		}
		fmt.Printf("Config dir created at %s\n", CONFIG_DIR)
	}

	if !utils.FileExist(CONFIG_FILE) {
		if err := os.WriteFile(CONFIG_FILE, []byte(CONFIG_EXEMPLE), 0644); err != nil {
			return err
		}
		fmt.Printf("Config file created at %s\n", CONFIG_FILE)
	}

	return nil
}

func InitEpubDir() error {
	epubDir := os.Getenv("EPUB_DIR")

	if !utils.FileExist(epubDir) {
		if err := os.MkdirAll(epubDir, 0755); err != nil {
			return err
		}
		fmt.Printf("Epub dir created at %s\n", epubDir)
	}

	return nil
}
