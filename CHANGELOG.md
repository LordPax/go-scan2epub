# Changelog

## [1.3.0] - 2024-08-10

### Added

* Add custom localize to translate text
* Add translation for command `base` `convert` `exist` `interval`
* Add translation for service `convert` `download` `scan-to-epub`

### Changed

* Move main command in `commands/base.go`
* Minor text changes
* Improve install script
* rename `service/scanToEpub.go` to `service/scan-to-epub.go`

## [1.2.0] - 2024-05-07

### Added

* Add log system
* Add command interval

### Changed

* move image processing from `utils/utils.go` to `utils/image.go`

## [1.1.0] - 2024-04-21

### Added

* Convert webp images to jpg

### Changed

* Don't stop process when a chapter fails to download
* Use `path.Join` instead of formatting path manually

## [1.0.0] - 2024-04-16

### Added

* Add command conv to convert pages to EPUB file
* Add command exist to check if a manga exists
* Add this changelog
* Download pages of manga
* Convert pages to EPUB file
* Rotate image when needed
