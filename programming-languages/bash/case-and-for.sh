# echo "Enter your age"
# read -r age
age=12

case "$age" in
  12) echo "$age"
  ;;
  18|19) echo 18 or 19
  ;;
  *) echo default
  ;;
esac

for Variable in $(seq 1 3)
do
  echo "$Variable"
done

for ((a=1; a <= 3; a++))
do
  echo $a
done





















































