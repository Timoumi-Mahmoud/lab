#!/bin/bash
# Script for adding new user, create home directory, set default /bin/bash login shell and add it to sudoers List.
#

read -p "Enter username: " username
read -p "Enter password: " password
useradd -m -s /bin/bash -p $(openssl passwd -1 $password) $username
if [ $? -eq 0 ]; then
		usermod -a -G sudo $username
		echo "$username ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers
		echo "User $username created successfully!"
		echo "User $username added to sudo group!"
else
		echo "Error while creating user!"
fi



