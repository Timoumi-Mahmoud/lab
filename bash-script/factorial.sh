#!/bin/bash
read -p "Enter number " num1

temp=1

for ((i=1; i<$num1; i++))
do
		temp=$((temp*i))
done
echo "The factorial of $num is: $temp"


