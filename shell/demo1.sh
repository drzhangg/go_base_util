url=http://c.biancheng.net/shell/
echo $url

name=哈哈哈
echo $name


lanuge=golang
echo "i use ${lanuge} to work"

unset name
echo $name

begin_time=`date`

sleep 10s

finish_time=$(date)
echo "begin time: ${begin_time}"
echo "finish time: ${finish_time}"

echo ${#lanuge}

