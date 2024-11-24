# Scan to epub

## Description

Golang project, which aims to download Scan chapters and convert them into EPUB for e-reader.

## Source

Scans can be downloaded from [lelscans](https://lelscans.net/lecture-ligne-one-piece)

## Source url format

```
http://www.example.com/{chap}/{page}.{ext}
```

* `{chap}`: Chapter number
* `{page}`: Page number
* `{ext}`: Image extension (png, jpg, jpeg, webp)
 
## Build and install

1. Clone the repository:

```bash
git clone https://github.com/LordPax/go-scan2epub.git
cd go-scan2epub
```

2. Build the project:

```bash
go mod download
go build
./install.sh
```

3. Execute the script to generate config

Will generate a config file at `~/.config/scan2epub/config.ini` create directory `~/scan2epub` to save converted epub

```bash
./scan2epub
```

4. Modify your config at `~/.config/scan2epub/config.ini`

## Config example

```ini
default="onepiece"

[onepiece]
name="One Piece"
author="Echiro Oda"
url="https://lelscans.net/mangas/one-piece/{chap}/{page}.{ext}"
epub_dir="/home/lordpax/scan2epub"
# epub_dir="/books" # Use this line if you are using docker
file_name="{author} - {name} {chap}.epub"
description="Scan of One Piece generated by scan2epub"

start_at=0
format=true
create_dir_per_file=true
cron="* * * * *"
cron_chap=1132
```

* `default`: Default manga to download
* `name`: Name of the manga
* `author`: Author of the manga
* `url`: URL format of the manga
* `epub_dir`: Directory to save the generated epub
* `file_name`: Name of the generated epub
* `description`: Description of the manga

* `start_at`: Chapter to start downloading
* `format`: Add "0" to page number if less than 10
* `create_dir_per_file`: Create a directory per chapter
* `cron`: Cron format to download the manga
* `cron_chap`: Chapter to start downloading when using cron

## Epub file name format

```
{author} - {name} {chap}.epub
```

* `{author}`: Author of the manga
* `{name}`: Name of the manga
* `{chap}`: Chapter number

## Docker

```bash
docker run -d --rm --name scan2epub -v ./config:/root/.config/scan2epub -v ./books:/books lordpax/scan2epub:latest
```

## Docker compose

```yaml
version: '3.7'

services:
  scan2epub:
    image: lordpax/scan2epub:latest
    container_name: scan2epub
    volumes:
      - ./config:/root/.config/scan2epub
      - ./books:/books
```
