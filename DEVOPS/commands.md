--------------------------------------------
### Commands

$ netstat -lpnt
- create user for every  application and give it only the permission it needs to run that app.
- Don't work with the Root user
$ adduser bara
$ usermod -aG sudo bara
$ su - bara
$ get the public ssh key , and place it in bara/home/.ssh/authorized_keys
$ ssh bara@server_ip_address

$ scp build/libs/java-example.jar root@IP:/root
$ java -jar java-example.jar &
$ ps aux | grep java
$ netstat -lpnt

$ useradd nexus
$ chown -R nexus:nexus nexus-3.28.1-01     //change ownership 

$ vim  nexus-3.28.1-01/bin/nexus.rc
$ /opt/nexus-3.28.1-01/bin/nexus start
$ cat /opt/sonatype-work/nexus3/admin.password



$ docker run -itd --name myAlpine  alpine:latest
$ docker run -itd -p host_port:container_port image
$ docker logs container_id|name


$docker exec -u 0 -it id_container bash //enter shell as root user


