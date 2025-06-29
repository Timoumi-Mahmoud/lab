#!/bin/bash
read -p "please enter your string :  " str1
echo $str1 | rev

echo ${#str1}
## To do check if the input is word or a sentence if sentese revervser the order of words , but if it's a single word reverse the word
