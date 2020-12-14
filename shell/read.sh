## !/bin/bash
#
#read -p "Enter some information > " name url age
#
#echo "网站名字：${name}"
#echo "url：${url}"
#echo "age: ${age}"

#read -n 1 -p "Enter a char>" char
#printf "\n"
#echo ${char}


#if
#  read -t 20 -sp "请在20秒内输入密码：" pass1 && printf "\n" &&
#  read -t 20 -sp "请在20秒内再次输入密码：" pass2 && printf "\n" &&
#  [ ${pass1} == ${pass2} ]
#then
#  echo "密码正确"
#else
#  echo "密码错误"
#fi

#read a
#read b
#
#(( ${a} == ${b} ))
#
#echo "exited status:"$?

read filename
read url

if test -w $filename && test -n $url;then
  echo $url > $filename
  echo "写入成功"
else
  "写入失败"
fi
