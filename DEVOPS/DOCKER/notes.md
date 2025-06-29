#### Docker

##### VM vs Container
- VM is a hardware abstruction: it takes physical CPUs and RAM from a host, and divides and shares it accross several VMS.
OS and application running inside the VM but the virtualization software usually has no real knowledge of that.
- Container is an application OS abstraction


- Container running instance of an image, each container is isolated from the other container even when created from the same image.

```
$ docker commit CONTAINER_ID
$ docker image tag IMAGE_ID imageName

$ docker build -it hello:v0.1 .
$ docker image inspect --format "{{ json .RootFS.Layers }}" 
```

- Layer : A Docker image is build up from a series of layers. Each layer represetnts an instruction in the image's dockerfile.
- Dockerfile: file container all the commands to create an image

```
FROM openjsk:8-jre-alpine
EXPOSE 6969
COPY ./build/libs/java-app-1.0-SNAPSHOT.jar /usr/app/
WORKDIR /usr/app
ENTRYPOINT ["java", "-jar", "java-app-1.0-SNAPSHOT.jar"]

```

- Docker Compose: is used to control muliple containers on a single system
- Docker Swarm: combines the ablitly to not only define the application architecture like compose, but to define and maintain high availabitity  levels, scaling, load balancing and more. 

```
$ Docker swarm init --advertise-add $(hostname -i)
$ docker swarm join-tokken manager // add master nodes
$ docker swarm jow --token <token> <host>

$ docker node ls
```

#### What is Container
- is a way to package application with all the necessary dependencies and configuration.
- Portable artifact
- Storage : Container repo public , or Private repo for companies 

Before containers: 
-> Dev team work install all library local dev env multiple steps of installation is error prone
-> dev produce an artifact then move it to server , but all additional services need to install on the server like database, external dependencies
After container:
-> no need to install any services because every container is own isolated environment  one command to install the app
-> no environmental configuration needed on the server

other container tech: containerD, cri-o, podman

#### Container vs Image
- container: layer of images mostly linux base image(alpine) | intermediate layer 
advantage: only different layers are downloaded
- image is a blueprint of a container 
- we can create many container from the same image
- Container is the running version of the image
#### Docker vs VM
OS: has two layers
Applications //layer 2
OS Kernel    //layer 1
Hardware

- Both are virtualization tools
- Docker virtualize the application layer
- VM virtualize the application layer and OS kernel
- # Size, speed, compatibility(can run on any OS host )


- After installing Docker we get :
1. Docker Engine: server responsible of pulling images, managing images and containers
		1. container Runtime : container lifesycle
		2. persistion data
		3. configuring network 
		4. build own Docker images
2. Docker Api : interacting with Docker server
3. Cli: client to execute docker commands


other container runtime: containerD, cri-o

#### Main Docker commands
see docs
#### Develop with container
Dev ---> commit (git)---> Jenkins ---> Jenkins CI produce artifact and build docker image from that artifact---->  push the image to Private Repo --> docker image deployed on pord server

## Demo Project:

[(3000)Node app (27017)]-----------> [mongoDB]<--------[MongoExpress(8081)]

- mongoDB and MongoExpress in the same network so they can communicate with each other without specifying the Ip address or the port name , I can use there name directly 
```bash
docker run -p 27017:27017 -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password --name mongodb --net mongo-network mongo

docker run -d -p 8081:8081 -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin -e ME_CONFIG_MONGODB_ADMINPASSWORD=password -e ME_CONFIG_MONGODB_SERVER=mongo --net mongo-network --name mongo-express mongo-express:latest

```

```yaml
version: '3'
services:
  monogdb: # container name
    image: mongo
    ports:
      - 27017:27017 # HOST:CONTAINER
    volumes:
      - db-data:/var/lib/mysql/data
    environment:
      - MONGO_INITDB_ROOT_USERNAME=admin 
      - MONGO_INITDB_ROOT_PASSWORD=password

  mongo-express:
    image: mongo-express:latest
    ports:
      - 8081:8081
    environment:
      - ME_CONFIG_MONGODB_ADMINUSERNAME=admin 
      - ME_CONFIG_MONGODB_ADMINPASSWORD=password 
      - ME_CONFIG_MONGODB_SERVER=mongodb

volumes:
  db-data

```

### Publish docker image on Nexus or ECR Amazon Elastic Container Registry or Docker registries(public)
1. login
2. Tag the image: registryDomain/imageName:tag (registryDcomain is docker.io/library) 
```bash
$ docker tag my-app:latest registryDomain/my-app:1.0
```
3. push 
```bash
$ docker push registryDomain/my-app:1.0
```

### Docker volumes:
- Allows data to persist and be shared 
- Folder in physical host file system is mounted into the virtual file system of Docker
Three types: 
1. Host Volumes
```bash
docker run -v /home/mount/data:/var/lib/mysql/data mysql
```
I decide where on the host file system the reference is made

2. Anonymous Volumes
```bash
docker run -v /var/lib/mysql/data mysql
```
/var/lib/docker/volumes/random-hash/data automatically created by Docker

3. Named Volumes
Improvement of Anonymous volumes
```bash
docker run -v name:/var/lib/mysql/data mysql
```
I can reference the volume by name.
should be used in production.


Docker volume location: /var/lib/docker/volumes


----------------------------------------------------------------------------------------------------------------------------------------
## Build Automation 
- Build app, build docker images , push to nexus 
dev --push-->Git repo --> should be tested and build
- publish new release: login to correct Docker repo, setup test environment to run tests, execute tests, build image and publish to repo

===>  Solution automate this process (detected server to automatise build process)

A detected server automatically will be trigger to run the build process , also will store all credentials: 
- Test code -> build app -> push to repo -> deploy to server
Tools: Jenkins

- What I can do with jenkins: Fucking middle man
1. Run Tests
2. Build artifacts
3. Publish artifacts
4. Deploy artifacts
5. Send notifications
+ many Plugins

```
docker run -p 8080:8080 -p 5000:5000 -d -v jenkins_home:/var/jenkins_home jenkins/jenkins:lts
```
- Two roles for Jenkins app
1. Adminstrator
administers and manages jenkins 
sets up jenkisn cluster
install plugins
backup jenkins data
2. Users
developers or devops teams
creating the actual jobs to run workflows


Jenkins job type:
1. Freestyle project
2. Pipeline                 X
3. Multi-configuration project
4. Folder
5. Github Organization
6. Multibranch Pipeline     X



## Example:
```yaml
FROM node:20-alpine

RUN mkdir -p /usr/app
COPY package*.json /usr/app/
COPY app/* /usr/app/
RUN npm install

CMD["node", "server.js"]
```

```yaml
FROM openjdk:8-jre-alpine
EXPOSE 8080

COPY ./build/libs/java-app-1.0-SNAPSHOT.jar /usr/app/
WORKDIR /usr/app

ENTRYPOINT ["java", "-jar" , "java-app-1.0-SNAPSHOT.jar"]
```

```
version: '3'
services:
  jenkinsServer:
    image: jenkins/jenkins:lts
    ports:
     - 8080:8080
     - 5000:5000
    environment:
     - Username: "mahmoud"
    volumes:
     - jenkins_home:/var/jenkins_home
volumes: 
    jenkins_home



```


