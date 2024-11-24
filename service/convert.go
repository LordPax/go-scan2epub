package service

import (
	"fmt"
	"os"
	"path"
	"scan2epub/config"
	"scan2epub/lang"
	"scan2epub/utils"
	"strings"

	epub "github.com/go-shiori/go-epub"
)

func convertChap(chap string) error {
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	log.Printf(l.Get("convert-chapter"), chap)

	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	epubDir := config.CONFIG_INI.Section(defaultSource).Key("epub_dir").String()
	createDir := config.CONFIG_INI.Section(defaultSource).Key("create_dir_per_file").MustBool()
	name := config.CONFIG_INI.Section(defaultSource).Key("name").String()
	tempDir := os.Getenv("TMP_DIR")
	pathChap := path.Join(tempDir, chap)

	if !utils.FileExist(pathChap) {
		return fmt.Errorf(l.Get("chapter-not-found"), chap)
	}

	if !utils.FileExist(epubDir) {
		if err := os.MkdirAll(epubDir, 0755); err != nil {
			return err
		}
	}

	if createDir {
		epubDir = path.Join(epubDir, fmt.Sprintf("%s %s", name, chap))
		if !utils.FileExist(epubDir) {
			if err := os.MkdirAll(epubDir, 0755); err != nil {
				return err
			}
		}
	}

	pages := getPageFromDir(pathChap)
	if len(pages) == 0 {
		return fmt.Errorf(l.Get("pages-not-found"), chap)
	}

	if err := createEpub(pages, epubDir, chap); err != nil {
		return err
	}

	return nil
}

func createEpub(pages []string, epubDir string, chap string) error {
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	defaultSource := config.CONFIG_INI.Section("").Key("default").String()
	author := config.CONFIG_INI.Section(defaultSource).Key("author").String()
	description := config.CONFIG_INI.Section(defaultSource).Key("description").String()
	name := config.CONFIG_INI.Section(defaultSource).Key("name").String()
	fileName := config.CONFIG_INI.Section(defaultSource).Key("file_name").String()

	epubFile, err := epub.NewEpub(name + " chapter " + chap)
	if err != nil {
		return err
	}

	epubFile.SetAuthor(author)
	epubFile.SetDescription(description)

	if err := addCover(epubFile, pages[0]); err != nil {
		return err
	}

	for _, page := range pages[1:] {
		if err := addImageToEpub(epubFile, page); err != nil {
			return err
		}
	}

	epubFileName := replaceValue(fileName, map[string]string{
		"{chap}":   chap,
		"{name}":   name,
		"{author}": author,
	})
	epubPath := path.Join(epubDir, epubFileName)
	if err := epubFile.Write(epubPath); err != nil {
		return err
	}

	log.Printf(l.Get("epub-created"), epubPath)
	return nil
}

func addCover(epubFile *epub.Epub, page string) error {
	if ext := path.Ext(page); ext == ".webp" {
		img, _, err := utils.DecodeImage(page)
		if err != nil {
			return err
		}

		dir, name := path.Split(page)
		fileName := strings.Split(name, ".")[0]
		page = path.Join(dir, fileName+".jpg")

		if err := utils.EncodeImage(page, img, "jpeg"); err != nil {
			return err
		}
	}

	cover, err := epubFile.AddImage(page, "")
	if err != nil {
		return err
	}

	if err := epubFile.SetCover(cover, ""); err != nil {
		return err
	}

	return nil
}

func addImageToEpub(epubFile *epub.Epub, page string) error {
	img, format, err := utils.DecodeImage(page)
	if err != nil {
		return err
	}

	width, height := utils.DimensionsImage(img)
	if width > height {
		img = utils.RotateImage(img)
	}

	if format == "webp" {
		dir, name := path.Split(page)
		fileName := strings.Split(name, ".")[0]
		page = path.Join(dir, fileName+".jpg")
		format = "jpeg"
	}

	if err := utils.EncodeImage(page, img, format); err != nil {
		return err
	}

	image, err := epubFile.AddImage(page, "")
	if err != nil {
		return err
	}

	body := "<img src=\"" + image + "\" width=\"91%\" />"
	_, err = epubFile.AddSection(body, "", "", "")
	if err != nil {
		return err
	}

	return nil
}
