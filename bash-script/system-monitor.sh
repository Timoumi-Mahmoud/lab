#!/bin/bash
# Bash script for monitoring CPU and Memory usage.
# User have the ability to set the duration of monitoring , in case no parameter is entered the default interval will be 10.
# Timoumi mahmoud 

time_interval=10
if [[ ! $1 ]]
then
		echo "no Time interval is set , default 10 seconds ! "
else
time_interval=$1
fi

while true
do
		cpu=$(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | awk '{print 100 - $1"%"}')
		mem=$(free -m | awk 'NR==2{printf "%.2f%%", $3*100/$2 }')
		echo "$(date) CPU Usage: $cpu, Memory Usage: $mem"
		sleep $time_interval
done
