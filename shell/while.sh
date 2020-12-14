i=1
sum=0

while  ((i<=100)) ; do
    ((sum+=i))
    ((i++))
    echo "i=${i}"
    echo "sum=${sum}"
done
echo "sum:${sum}"
