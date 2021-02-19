#!/usr/bin/env bash

# e.g. init.sh {DIR} {LINK}

D=$1
LINK=$2

if [ -d "$D" ]; then
	echo "$D exists"
	exit 0
fi

echo "your input, dir: $D, link: $LINK"
mkdir ${D}

cd ${D} && go mod init main && touch main.go && echo -e "package main // $LINK\n\nfunc main() {}" > main.go