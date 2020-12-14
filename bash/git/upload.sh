#/bin/sh

cd D:/godemo/go_base_util

pwd

git add .

git config --global user.name "drzhangg"

git config --global user.email "654014730@qq.com"

git commit -am "$1"

git push
