# Scan to epub

## Description

Golang adaptation of [scan-to-epub](https://github.com/LordPax/scan-to-epub.git) project, which aims to download Scan chapters and convert them into EPUB for my e-reader.

## Source

* Scans can be downloaded from [lelscans](https://lelscans.net/lecture-ligne-one-piece)

## Source url format

```
http://www.example.com/{chap}/{page}.{png|jpg|jpeg|webp}
```

## Pre-requis

* Converts .webp files to other image formats and vice versa [libwebp-1.1.0](https://developers.google.com/speed/webp/docs/compiling)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/LordPax/go-scan2epub.git
cd go-scan2epub
```

2. Build the project:

```bash
go mod tidy
go mod vendor
go build
```

3. Execute the script to generate config

Will generate a config file at `~/.config/scan2epub/config` create directory `~/scan2epub` to save converted epub

```bash
./scan2epub
```

4. Modify your config at `~/.config/scan2epub/config`:

**Example**

```bash
URL=https://lelscans.net/mangas/one-piece
EPUB_DIR=/home/username/scan2epub
```
