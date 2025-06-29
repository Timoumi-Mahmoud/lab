#!/bin/bash
read -p "Please enter your Email: " email

if [[ $email =~ ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$ ]]
then
		echo "This is a valid email !"
else
		echo "This is not a valid email !"

fi
