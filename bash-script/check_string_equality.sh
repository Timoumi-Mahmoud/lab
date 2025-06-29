#!/bin/bash
read -p "please enter the first string :  " str1
echo $str1
read -p "please enter the second string :  " str2
echo $str2

if [ $str1 == $str2 ]; then
		echo "$str1 and $str2 are equal"
else
		echo "$str1 and $str2 are not equal"
fi

