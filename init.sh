#!/usr/bin/env bash

# e.g. init.sh {DIR} {LINK}

ID=$1
TITLE=$2
LINK=$3

DIR=$ID.$TITLE

if [ -d "$DIR" ]; then
	echo -e "\033[31m 目录${DIR}已经存在 \033[0m"
	exit 0
fi

echo "创建目录${DIR}, 并开始初始化"
mkdir ${DIR}

cd ${DIR} && go mod init main && touch main.go && echo -e "package main // $LINK\n\nfunc main() {}" > main.go
