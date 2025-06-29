#!/bin/bash
let blah=99999999999
echo $blah


my_array=("bara" "khdija" "mahmoud" "salim")

echo ${my_array[@]}
my_array[8]="non non"
echo ${my_array[@]}
echo ${my_array[2]}
echo "length: ${#my_array[@]}"

my_array=("new value" ${my_array[@]} ) # add at the beginning 
my_array=( ${my_array[@]} "last value" ) # add at the end

unset my_array[2]
echo ${my_array[@]}
echo -e "------------------------\n"

#for arg in `seq 1 10`
#do
#		echo $arg
#		rm test.$arg
#done



read counter
while [ $counter -ge 0 ]
do let counter--
		echo $counter
done

read counterTwo
until [ $counterTwo -lt 0 ]
do let counterTwo--
		echo $counterTwo
done



#my_array=(first ${my_array[@]})


#echo $#
#sudo find . -type f -name '*.log' | while read fname ; do echo mv  ; done | sh -x
