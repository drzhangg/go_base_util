nums=([3]=24 [4]=55 [6]="jerry" [10]="http://www.bilibili.com")
echo "nums:"${nums[6]}
echo "nums:"${nums[3]}
echo "nums:"${nums[*]}
echo "nums:"${nums[@]}

echo "length:"${#nums[*]}

unset nums[1]
unset nums[3]
echo "len:"${#nums[*]}
echo "nums:"${nums[3]}

num2=("i want a job" "salary 16000")
echo ${num2[@]}

new_arr=(${nums[@]} ${num2[*]})
echo ${new_arr[@]}
