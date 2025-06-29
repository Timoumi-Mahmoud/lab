#!/bin/bash
read -p "please enter your string :  " str1
echo $str1
read -p "Do you want want to convert it to Upper or Lower : (U|L)" choise
#echo $choise

if  [ ! -n "$choise" ]; then
		echo "This field cannot be empty"
		exit 1
elif [ $choise == "U" ]; then
		echo $str1 | tr '[:lower:]' '[:upper:]'
elif  [ $choise == "L" ]; then
		echo $str1 | tr '[:upper:]' '[:lower:]'
else
		echo "Please chose U for Upper or L for Lower"
fi

