#!/usr/bin/env bash

# 这段脚本是用来对GOPATH项目编译前的处理, Go module项目无需关心GOPATH
# cpath=`pwd`
# PROJECT_PATH=${cpath%resk.com*}
# echo $PROJECT_PATH
# export GOPATH=$GOPATH:${PROJECT_PATH}

# 编译
SOURCE_FILE_NAME=main #源文件名称
TARGET_FILE_NAME=reskd #编译后目标文件名称

rm -fr ${TARGET_FILE_NAME}* #首先把之前打好的包删除掉

build(){ #构建脚本
	echo $GOOD $GOARCH
    tname=${TARGET_FILE_NAME}_${GOOS}_${GOARCH}${EXT}
	#编译
	env GOOS=${GOOS} GOARCH=${GOARCH} \
	go build -o ${tname} \
	-v ${SOURCE_FILE_NAME}.go
	#添加可执行权限
	chmod +x ${tname}
	#打包
	mv ${tname} ${TARGET_FILE_NAME}${EXT}
    if [ ${GOOS} == "windows" ];then
        zip ${tname}.zip ${TARGET_FILE_NAME}${EXT} config.ini *.sh ../public/
    else
        tar --exclude=*.gz  --exclude=*.zip  --exclude=*.git -czvf ${tname}.tar.gz ${TARGET_FILE_NAME}${EXT} config.ini *.sh ../public/ -C ./ .
    fi
    mv ${TARGET_FILE_NAME}${EXT} ${tname} #打包完再换回来
}

CGO_ENABLED=0

# mac osx
GOOS=darwin # darwin:mac_osx linux:linux windows:windows
GOARCH=amd64 # amd64 架构
build # 执行编译

#linux
GOOS=linux
GOARCH=amd64
build

#windows
GOOS=windows
GOARCH=amd64
build