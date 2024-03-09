#!/usr/bin/env bash

echo "Hello world! This is bash scripting"

echo "Message 1"; echo "Message 2"

myName="Gabriel Rockson"
echo "$myName"

echo "${myName/Gabriel/Kweku}"

length=7
echo "${myName:0:7}"

echo "${myName: -3}"

echo "${#myName}" # this should return the length of the string

otherVariable="myName"
echo ${!otherVariable}

foo="Default Value Available"
echo "${foo:-"Default Value IF Foo Is Missing Or Empty"}"


array=(one two three four five six)
echo "$array"
echo "${array[0]}"
echo "${array[1]}"

echo "${array[@]}"
echo "${array[@]:0:6}"

# For loop to print all the elements each of them on a new line
for number in "${array[@]}";
do
  echo "$number"
done

# Built-in variables
# To get last return Value -> $?
# To get script's PID -> $$
# To get Number of arguments passed to script -> $#
# To get All arguments passed to script: $@
# To get Script's arguments separated into different variables: $1 $2...
#

# Brace expansion
echo {1...10}
echo {a...z}

echo "$PWD"
echo "I'm in $(pwd)"
echo "I'm in $PWD"

# Reading a value from input
echo "What's your name?"
read -r name
echo "Hello, $name!"
echo "I'd like to know your age: "
read -r age
echo "Wow, thanks for saying your age $age"















