#
## 
- Container are isolated environments.
- One computer/server can run lots of containers, which share the same OS, CPU and memory.


Docker is a platform for running applications in lightweight units called container.

- How to get apps to the cloud ?
1. Iaas: every component runs in its own VM in isolation. It's an easier migration, but VMs are likely to be underutilized and expensive to run.

2. Paas: Every service get mapped to a managed service from the cloud provider.

3. Docker: migrate each part of my application to a container and then I can run the whole application in containers using aks, amazon , or on my own docker cluster in the datacenter.


* Serverless: developers write function code, put it to a service, and that service builds and packages the code. 
when the consumers use the function, there service starts an instance of the function to process the request.
But functions in the cloud aren't portable you can’t take your AWS Lambda function and run it in Azure. (lock-in)


### Architecture
`Docker engine`: is the management component of Docker, it looks after the local image cache, download images when i need them and reusing them if they're already downloaded.
It also works with OS to create container, virtual networks, and all the other docker resources. 
Is a background process that always running.
Makes all the features available through the Docker API,
which is just a standard HTTP-based REST API. You can configure the Engine to make the API accessible only from the local computer (which is the default), or make it available to other computers on your network.
 The Docker
Engine uses a component called containerd to actually manage containers, and containerd in turn makes use of operating system features to create the virtual environment that is the container.


`The Docker command-line interface (CLI) `is a client of the Docker API. When you run Docker commands, the CLI actually sends them to the Docker API, and the Docker Engine does the work

## Build share and run
1. package the app 
2. publish it to a public/private repository
3. any one with access can run the app in a container

## Container
---------------------
The machine name of every container will change every time, and the IP address will often change, but every container is running on the same computer, so where do these different machine names and network addresses come from? 

A Docker container is the same idea as a physical container think of it like a box with an application in it. Inside the box, the application seems to have a computer all to itself: it has its own machine name and IP address, and it also has its own disk drive (Windows containers have their own Windows Registry too). 
Those things are all virtual resources the hostname, IP address, and filesystem are created by Docker. They’re logical objects that are managed by Docker, and they’re all joined together to create an environment where an application can run. That’s the “box” of the container. The application inside the box can’t see anything outside the box, but the box is running on a computer, and that computer can also be running lots of other boxes.
The applications in those boxes have their own separate environments (managed by Docker), but they all share the CPU and memory of the computer, and they all share the computer’s operating system. 

 fixes two conflicting problems in computing: isolation and density. Density means running as many applications on your computers as possible,
 to utilize all the processor and memory that you have. But apps may not work nicely  with other apps they might use different versions of Java or .NET, they may use incompatible versions of tools or libraries, or one might have a heavy workload and starve the others of processing power. Applications really need to be isolated from each other, and that stops you running lots of them on a single computer, so you don’t get density.

VM:
- they give you a box to run my app in , but the box for a vm needs to contain its own operating system, it doesn't share the os of the computer where the vm is running.
[HOST machine]
[hypervisor]
[vm1] [vm2] [vm3]

*diff vm and container*
- size lightweight , cost(licensing cost)., start quickly and run lean
- porvide isolation at the cost of density.


docker ls -all which shows all containers in any status.containers are running only while the application inside the container is running. As soon as the application process ends, the container goes into the exited state.
Exited containers don’t use any CPU time or memory.

Containers don’t disappear when they exit. Containers in the exited state still exist, which means you can start them again, check the logs, and copy files to and from the container’s filesystem.





When you install Docker, it injects itself into your computer’s networking layer. Traffic coming into your computer can be intercepted by Docker, and then Docker can send that traffic into a container.
containers aren't exposed to the outside world by default. Each has its own IP address, but that’s an IP address that Docker creates for a network that Docker manages—the container is not attached to the physical network of the computer.
Publishing a container port means Docker listens for network traffic on the computer port, and then sends it into the container.In the nginx (see commands section), traffic sent to the computer on port 8088 will get sent into the container on port 80.

Docker containers also have environment variables, but instead of coming from the computer’s operating system, they’re set up by Docker in the same way that Docker creates a hostname and IP address for the container.




## Image
--------------------------------------------------------
- Docker image is logically one thing, you can think of it as big zip file that contains the whole application stack.
for example (node runtime + application code) 
During the pull you don’t see one single file downloaded; you see lots of downloads in progress. Those are called image layers. A Docker image is physically stored as lots of small files, and Docker assembles them together to create the container’s filesystem. When all the layers ave been pulled, the full image is available to use.


Docker images may be packaged with a default set of configuration values for the application, but you should be able to provide different configuration settings when you run a container.



### image layer
```bash
$ docker image history web-ping"

```
Each line in the Dockerfile creates an image layer.
A docker image is a logical collection of image layers. Layers ate the files that are physically stored in the Docker Engine cache.
Image layer can be shared between different images and different containers.
If image layers are shared around, they can’t be edited otherwise a change in one image would cascade to all the other images that share the changed layer. Docker enforces that by making image layers read-only.

### Dockerfile usage 
- Every developer will have build toolset to install like the compiler, linker, package manager, App runtime, which will take a bit of time to install in developer laptop and can be cumbersome (oh it works on my machine).
would be much cleaner to package the build toolset once and share it, which is exactly what you can do with Docker. You can write a Dockerfile that scripts the deployment of all your tools, and build that into an image. Then you can use that image in your application Dockerfiles to compile the source code, and the final out- put is your packaged application


- Multi-stage  Dockerfile , because there are several stages to the build.(several FROM commands)  ==> make images small in size and fast to luanch.

- JAVA: 
    stage 1: download dependencies and builds the application.
        from maven:jdk (large image)
        run mvn install
        run mvn package
    stage 2: copies the built application folder and run tests
        COPY --from=builder ./../iotd-service-0.1.0.jar .
        run mvn test
    stage 3: the final stage copies the tested application 
        COPY --from=builder ./../iotd-service-0.1.0.jar .
        ENTRYPOINT ["java", "-jar", "/app/iotd-service-0.1.0.jar"]

--> every stage has it's own cash 



(IF you want container to talk to each other, they have to be on the same docker network)

* Image reference:
[docker.io]/[timoumi-mahmoud]/[golang]:[latest]
(1) the registry domain, where the image is stored.
(2) the account of the image owner, cloud be an individual user or an org.
(3) the image repository name, one repository can store many version of an image.
(4) the image tag, used for versioning or variations of an application. Latest is the default.

```bash
$ export dockerId="<your-docker-id-goes-here>"
$ docker login --username $dockerId

$ docker image tagn image-gallery $dockerId/image-gallary:v1
$ docker image ls --filter reference=image-gallery --filter reference='*/image-gallery'

$ docker image push $dockerId/image-gallary:v1
```
---------------------------------------------------------
## Volumes:
stateful apps:
a docker container has a filesystem with a single disk drive, and the contents of that drive are populated with the files from the image.

*image layers are shared so they have to be read-only, and there is one writeable layer per container, which has the same life cycle as the container.
image layers have their own life cycle any images you pull will stay in your local cache until you remove them. But the container writeable layer is created by Docker when the container is started, and it’s deleted when conatainer is removed.(stopping the container doesn't automatically remove it)
The writable layer isn't just creating new files. A container can edit existing files form the image layers. But image layers are red-only so docker does 'copy-on-write' porcess to allow edits to files come from read-only layers. When the container tries to edit a file in an image layer, Docker actually makes a copy of that file into the writable layer, and the edits happen there. It’s all seamless for the container and the application, but it’s the cornerstone of Docker’s super-efficient use of storage.


* Docker cp:
        docker cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-
        docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH


Docker container volume is unit of storage, you can think of it as USB stick for containers.
Volumes exist  independently of containers and have their own life cycles, but they can attached to containers.

There is two ways to use Volumes with containers:
- inside the docker file : Volume /data
``` --volume-from container1 ``` (batch jobs database backup)
``` docker run -d -p 800:80 -v todo-list:/data -name blablabla ```
- mount: makes a directory on the host machine available as path on a container.

``` docker run --mount type=bind, source=$source,target=$target -d -p 800:80 -name blablabla ```


(-) When you mount a target that already has data, the source directory replaces the target directory so the original files from the image are not available.


--------------------------------------------------------------------
### Docker compose:
docker compose file describe the desired state of my app.(declarative)





















        


---------------------------------------------------------
## commands:
```go
$ docker container rm -f $(docker container ls -aq)
$ ocker container rm --force $(docker container ls --all --quiet)
$ docker image rm -f $(docker image ls -f reference='diamol/*' -q)
$ docker system prune


$ docker container top <id> # list the processes running in the container
$ docker container logs <id> 
$ docker container inspect <id> # shows all the details of a container

$  docker run --name fend  --detach --publish 8088:80 nginx
# --detach -d: start a container in the background and show it's id.
# --publish -p: publishes a port from the container to the computer

$ docker container stats <id>
CONTAINER ID   NAME         CPU %     MEM USAGE / LIMIT     MEM %     NET I/O         BLOCK I/O   PIDS
e2f33f0f3394   postgres15   0.00%     25.39MiB / 6.032GiB   0.41%     1.84kB / 126B   0B / 0B     6


$ docker container run --env TARGET=google.com diamol/ch03-web-ping

$ docker run --name fend  --detach --publish 8088:80 nginx
$ docker system df
```
```Dockerfile
FROM alpine AS build-stage
RUN echo 'Building...' >> /build.txt

FROM build-stage AS test-stage
COPY --from=build-stage /build.txt /build.txt
RUN echo 'Testing...' >> /build.txt

FROM alpine
COPY --from=test-stage /build.txt /build.txt
CMD ["cat", "/build.txt"]

# individual stages are isolated.
```

```Dockerfile
FROM diamol/maven AS builder
WORKDIR /usr/src/iotd
COPY pom.xml .
RUN mvn -B dependency:go-offline
COPY . .
RUN mvn package

# app
FROM diamol/openjdk
WORKDIR /app
COPY --from=builder /usr/src/iotd/target/iotd-service-0.1.0.jar .
EXPOSE 80
ENTRYPOINT ["java", "-jar", "/app/iotd-service-0.1.0.jar"]
```
