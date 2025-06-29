#!/bin/bash
read -p "first number " num1
read -p "second number " num2
echo "num1: $num1, num2: $num2"

read -p "Operation: + | - | / | % " operation
result=$(($num1 $operation $num2))
echo " $num1  $operation  $num2 = $result"

### Print even and odd number (of the result)
for (( i=1; i<=$result; i++ ))
do
		if [ $((i%2)) == 0 ]
		then
				echo "$i is even"
		else
				echo "$i is odd"
		fi
done

#### Print the muliplication table for result if it's < 10
if [ $result -lt 10 ]
		then
				for((j=0; j<10; j++))
				do
						echo "$j * $result : $(($result * $j))"
				done
		fi
### Print Random number between 1 and the value of result
echo $((1+ RANDOM % $result))

