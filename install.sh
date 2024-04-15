#!/bin/bash

[ ! -f scan2epub ] && echo "scan2epub not found" && exit 1
install -Dm 755 scan2epub ~/.local/bin/scan2epub
