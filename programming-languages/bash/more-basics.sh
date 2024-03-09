echo "What is your name?"
read name

# echo "$name"

# If structures in bash
if [ "$name" = "Gabriel" ]; then
  echo "Your name isn't your username"
else
  echo "Please enter the valid username"
fi

echo "Enter your age"
read -r age

if [ "$age" -eq 15 ] && [ "$name" = "Gabe" ]; then
  echo "Hey $name, you're $age, you're old enough to drink"
elif [ "$age" -eq 92 ]; then
  echo "You're tooo old"
else
  echo "Unknown behavior"
fi

# email=me@example.com
# if [[ "$email" =~ [a-z]+@[a-z]{2,}\.(com|net|org) ]]; then
#   echo "Valid email"
# fi
















