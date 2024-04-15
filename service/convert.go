package service

import (
	"fmt"
	"os"
	"scan2epub/utils"

	epub "github.com/go-shiori/go-epub"
)

func convertChap(chap string) error {
	fmt.Printf("Converting chapter %s to epub\n", chap)

	tempDir := os.Getenv("TMP_DIR")
	epubDir := os.Getenv("EPUB_DIR")
	pathChap := tempDir + "/" + chap

	if !utils.FileExist(pathChap) {
		return fmt.Errorf("Chapter %s not found", chap)
	}

	if !utils.FileExist(epubDir) {
		if err := os.MkdirAll(epubDir, 0755); err != nil {
			return err
		}
	}

	pages := getPageFromDir(pathChap)
	if len(pages) == 0 {
		return fmt.Errorf("No pages found for chapter %s", chap)
	}

	if err := createEpub(pages, epubDir, chap); err != nil {
		return err
	}

	return nil
}

func createEpub(pages []string, epubDir string, chap string) error {
	author := os.Getenv("AUTHOR")
	description := os.Getenv("DESCRIPTION")

	epubFile, err := epub.NewEpub("Chapter " + chap)
	if err != nil {
		return err
	}

	epubFile.SetAuthor(author)
	epubFile.SetDescription(description)

	cover, err := epubFile.AddImage(pages[0], "")
	if err != nil {
		return err
	}

	if err := epubFile.SetCover(cover, ""); err != nil {
		return err
	}

	for _, page := range pages[1:] {
		if err := addImageToEpub(epubFile, page); err != nil {
			return err
		}
	}

	epubFileName := fmt.Sprintf("%s/chap-%s.epub", epubDir, chap)
	if err := epubFile.Write(epubFileName); err != nil {
		return err
	}

	fmt.Printf("Epub created at %s\n", epubFileName)
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
		if err := utils.EncodeImage(page, img, format); err != nil {
			return err
		}
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
