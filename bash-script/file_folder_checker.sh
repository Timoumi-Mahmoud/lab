#!/bin/bash
echo $1

if [[ ! $1 ]]
then
		echo "Please enter your file/folder name ! "
		exit 1

fi

if [[ -f $1 ]]
then
		echo "$1 is a file"
		echo "here is the stat:$(stat $1)"
elif [[ -d $1 ]]
then
		echo "$1 is a folder"
		echo "here is the stat:$(stat $1)"
else 
		echo "No such file or folder !!"
fi

