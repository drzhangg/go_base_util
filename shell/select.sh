echo "what is your favourite OS?"

select name in "Linux" "Windows" "Mac OS" "UNIX" "Android";
do
  case $name in
    "Linux")
    echo ${name}
done
echo "you have selected ${name}"
