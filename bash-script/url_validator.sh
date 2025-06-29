#!/bin/bash
read -p "Please enter an URL: " url

if [[ $url =~ ^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$ ]]
then
		echo "This is a valid email !"
else
		echo "This is not a valid email !"

fi
