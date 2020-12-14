#function show(){
#  echo "Tutorial: ${a}"
#  echo "URL: ${b}"
#  echo "Author: ${c}"
#  echo "Total $# parameters"
#}
#
#read a b c
#
#show

function getsum() {
    local sum=0

    for n in $@
    do
      ((sum+=n))
    done

    echo ${sum}
    return 0
}

total=$(getsum 1 2 3 4 5)
echo ${total}

echo $(getsum 1 2 3)
