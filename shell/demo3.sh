function func() {
    echo "name: $1"
    echo "age: ${2}"
}

func jerry 24

function getSum() {
    local sum=()
    for n in $@
    do
      ((sum+=n))

    done

    return $sum
}

getSum 10 20 55 15
echo $?
