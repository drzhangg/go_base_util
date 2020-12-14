sum=0

#for ((i=1;i<=100;i++))
#do
#  ((sum+=i))
#  done
#echo "this num is:${sum}"

#for filename in $(ls *.sh)
#do
#  echo ${filename}
#done

function func() {
    for str in $@
    do
      echo $str
    done
}

func a d c b e
