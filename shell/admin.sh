#!/bin/bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

SERVER="chuanghehui_business"
LOG_PATH=./
DATE=$(date "+%Y%m%d")
LOG_FILE=$LOG_PATH/"$DATE""_business.log"
INTERVAL=1

# 命令行参数，需要手动指定
ARGS=""

function dockerMake() {
    go get github.com/jessevdk/go-assets-builder
    go-assets-builder view/templates -o assets.go
    docker rm app -f
    docker rmi app
    docker build -t app .
    docker run --name app -p 8082:8082 -d app
}

function start() {
        if [ "`pgrep $SERVER -f`" != "" ];then
                echo "$SERVER already running"
                exit 1
        fi

        git checkout .
        git pull

        go install
        # go-assets-builder view/templates -o assets.go
        nohup $GOPATH/bin/chuanghehui_business &>> $LOG_FILE &

  sleep $INTERVAL

        # check status
        if [ "`pgrep $SERVER -f`" == "" ];then
                echo "$SERVER start failed"
                exit 1
        fi
        echo "$SERVER: started, pid=$!"
}

function status() {
        if [ "`pgrep $SERVER -f`" != "" ];then
                echo $SERVER is running
        else
                echo $SERVER is not running
        fi
}

function stop() {
        if [ "`pgrep $SERVER -f`" != "" ];then
                kill -2 `pgrep $SERVER -f`
        fi

        sleep $INTERVAL

        if [ "`pgrep $SERVER -f`" != "" ];then
                echo "$SERVER stop failed"
                exit 1
        fi

        echo "$SERVER: stopped"
}

function restart() {
    stop
    sleep 1
    start
}

function help() {
    echo "start|stop|restart|status"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "stop" ];then
    stop
elif [ "$1" == "start" ];then
    start
elif [ "$1" == "restart" ];then
    restart
elif [ "$1" == "status" ];then
    status
else
    help
fi
