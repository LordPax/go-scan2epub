#!/bin/bash

[ "$EUID" -ne 0 ] && echo "Please run as root" && exit 1

function installFunc() {
	[ ! -f scan2epub ] && echo "scan2epub not found" && exit 1
	echo "Installing scan2epub to /usr/bin/scan2epub"
	install -Dm 755 scan2epub /usr/bin/scan2epub
}

function uninstall() {
	echo "Uninstalling scan2epub from /usr/bin/scan2epub"
	rm -f /usr/bin/scan2epub
}

case "$1" in
	help) echo "Usage: $0 {install|uninstall}" ;;
	uninstall) uninstall ;;
	*) installFunc ;;
esac
