#!/bin/bash

read -p "Enter two value(True|False)" val1

read -p "Enter two value(True|False)" val2

read -p "Entern an operation(and|or|not):  " op

case $op in
		and)
				if [[ $val1 == true && $val2 == true ]] ;then
						echo "Result: true"
				else
						
						echo "Result: false"
				fi;;

		or)
				if [[ $val1 == false && $val2 == false ]] 
				then
						echo "Result: false"
				else
						echo "Result: true"
				fi;;

		not)
				if [[ $val1 == true ]] 
				then
						echo "Result: false"
				else
						
						echo "Result: true"
				fi;;

		*) echo "Invalide oprator"
esac


