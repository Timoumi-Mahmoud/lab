#!/bin/bash
read -p "please enter your string :  " str1
echo $str1

str=${str1// /}
echo "Result: $str"

read -p "By wish character you want to replace white space: " flag
echo $str1 | tr {{" "}} {{$flag}}


