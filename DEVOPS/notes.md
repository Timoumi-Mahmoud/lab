# DevOps:
- Complete software Development Process:

Programming => Testing => Release : build && run on server

- Dev vs operations(run software in production) no downtime
- Problem with old approach : releases takes long

---------------------------------------
#### DevOps:
- Was just a way of working between DEV's and OP's
- Intermediate between dev's and op's 
- Implement -> Test -> Build -> Deploy -> Maintain

- Waterfall vs Agile:
plan everything beforehand: Requirements
developers code complete app: Development
testing after everything has developed: Testing
huge preparation: Operation
=> ineffective process
=> over time new requirements may arise
=> many places of failure and miscommunication

Agile 
=> speed of development, testing and deployment cycles
=> each feature gets tested, deployed
=> immediate feedback
=> fast development and deployment process
=> scrum and kanban
The hart of CI/CD
- Intersection of Dev & ops


----------------------------------------------------------------------------------------------------------------------------------------

### Version Control with Git
- alternative SVN 
- Remote Git Repo, cloud repo: where the code lives github, gitlab, or Companies own Git servers like Bitbucket, gitlab(localy)
- Local Git Repository : local copy of the code
- History: of code changes (git log)
- Staging: changes to commit


Working Directory      staging Area      local Repo    Remote Rep
			---git add-->      ---git commit-->   ---git push-->
													<---git pull---



- .git : contains information about the repository (History, log, config extra)

- Use ssh for clonning repo without the need to provide username and password: ssh-keygen -t rsa -b 4096 -C "email@example.com" 
		  copy the rsa.pub

##### Best practises 
1. Branches: 
- dev branch: intermediary master
- master branch : ready for production 
- other branches: fix_bug , features
=> during the sprint: merge features and bugfixes into dev.
=> at the end of the sprint: dev branch gets merged with master branch.

2. Merge Requests:
- other developer reviews code changes before merging , big feature or junior developer or expertise in different area.
- if In my local branch for example : bug_fix , I am missing 2 commits from the master branch :
	pull changes from master branch
	merge it with git merge master

3. Deleting the branches after the merge
- Leave the branch or delete it
$ git checkout master
$ git pull
$ git branch -d feature/database-connection
$  git push origin --delete feature/database-connection

4. Ignoring files /directory 
$ add file and directories in .gitignore file
$ git rm -r --cached directories_that_was_already_tracked
$ commit changes

5. Stash
$ git stash
$ git stash pop or apply

6. Undoing and changing commits
$ git reset --hard HEAD~1 // ~Number of commits i want to revert to  and discard changes
$ git reset HEAD~1  // undo the commits but keep the changes(soft)
$ git commit --amend  // new commit but will be included with last commit
- In case I pushed the commit to the remote repo: 
$ git reset --hard HEAD~1 
$ git push --force
- Or Use git revert:
git revert commit-hash // it will create new commit 
! Don't do this in master and develop branch 
only when working alone in a branch 


##### Git for DevOps
- Used in Infrastructure as Code: many kubernetes/Terraform/Ansible config files
- Bash/Python Script 
==> Tracked - history of changes
==> Securely stored in one place
==> shareable for DevOps team
- CI/CD Pipeline and Build Automation 
checkout code, test and build application, etc
need integration for the build automation tool with app git repository 
I need to setup integration with build automation tool and repo
for example: getting commit hash of specific commit or 
check if changes happen in frontend or backend code.

----------------------------------------------------------------------------------------------------------------------------------------------

### Database
- Apps use Database to persist data
1. option 1
	each developer install mysql db localy , each dev has own DB with own test data
=> can't mess up someone else's test DB data.
=> start from empty DB , test operation crud
=> installation and setup for every dev
2. option 2
	DB hosted remotely 
=> no need for local installation 	
=> Test data available right from the start
=> Impact other dev data 

=====> Ideal solution : have both

[PROD] 
[Staging]
[Dev]

Tell the library (db connection) which db to connect to. (change credentials and endpoint)=> good practice is to use Env variable
Dev, test , prod Database each with different db connect: env variable.from the cli or code editor.
For example in JAVA : the app.yml file :
							application-dev.yml
							application-prod.yml
							application-staging.yml
in node js: config.js
- Logging level: log more in dev mode, log less in prod mode (only errors and warnings)
- DB endpoint/credentials
- Third party API's 

###### Database best practices:
- replicate
- regular backups
- performs under high load
- restore DB

###### DB types:
1. Key value Database
- Redis, Memcached, etcd from kubernetes
- unique key, no joins
- very fast : store in memory not like other db which store data in the hard drive
- Limited storage, not used as primary DB best use case is for caching, or for Message Queue(message broker)

2. Wide Column DB
- Cassandra, Apache HBase
- Same key but data is divided into many column
- Doesn't have a predefined schema : schema-less
- key : columns
- no joins, Queries similar to SQL , very scalable, not used as primary DB
- best for handling large amounts of unstructured data like time-series, iot-Records or Historical-records

3. Document DB
- MangoDB, DynamoDB, CouchDB
- Document for key value pair , schema-less, no joins
- Documents are grouped in Collections  and collections can be organize in relational hierarchy 
- Slower in writes, fast to read
- Best for Mobile apps, Game apps, CMS
- can be used as primary DB
- not used for graphs

4. Relational DB
- Mysql, PostgreSql
- most used, schema and data types needs to be created first
- Structured Query Language SQL
- Tables: columns and rows
- Normalizing to avoid duplicated data
- ACID: Atomicity, Consistency, Isolation, Durability 
Either all changes get applied or NON, example: transaction on 10 tables , network connectivity disruption in the middle of transaction.
No half-changes are updated in database
- Difficult to scale (but CockroachDB comes to rescue)

5. Graph DB
- Neo4j, Dgraph 
- like reddit , user can subscribe to many subreddit  or Youtube( sub user can scribe to many channels) 
- Directly connect Entities : nodes and edges, edges are the relationships.
- Easy to query 
- Best for Graphs, Patterns, Recommendation engine

6. Search DB
- ElasticSearch , Solr
- Search through massive data entries 
- Full text search in efficient and fast way 
- Similar to document-orientated database, the difference is in the search DB in background, will analyse all text from data object and it will create index of all individual words or data pieces, index is searched only 

==> I could combine many database for example: 
For most of data => Relational DB
To handle fast search => Search Engine DB
Cache => Key Value DB

----------------------------------------------------------------------------------------------------------------------------------------------

## Build Tool & Package Manager
- Technologies used in software development process and as part of the application building process 
- Npm, Gardle, Maven

- After dev -> app needs to be deployed on a production server, move the code and it's dependencies.
- Package app into a single movable file==> calling artifact 
- Packaging ="building the code"
- Involve : compiling, compress from 100s of files into one file => Artifact
- After that keep artifact in storage to deploy it multiple times(dev, test, production, have backup
- Artifact Repository : Nexus, JFrog Artifactory
- Artifact file : java (war, jar) whole code plus dependencies  

* Build Tool: Install dependencies,Run tests, Compile and compress my code
##### Java
- JAR (Java Archive) , WAR
- Maven uses XML , Gradle uses Groovy language
- Both have a command line tool and commands to execute the tasks 
- build.gradle and pom.xml(add a plugin)

$ ./gradlew build  => generate build folder: /libs java.app-1.0-Snpshot.jar
$ mvn install => /target


$ java -jar name_jar_file

##### JavaScript
- no special artifact type 
- npm or yarn: are package managers not build tools
- package.json file for dependencies 
- webpack: transpiles | minifies | bundles | compresses the code

$ npm start
$ npm stop
$ npm test
$ npm publish
$ npm pack : create a tar file if extracted I will find the application under /package and source code(without node_modules)

$ npm run build // server.bundle.js

** For Front end for example React : code needs to be transpiled in order to webbrowser to understanded it, also needs to be compressed/minified so that does'nt be too large for browser to donwnload.
$ npm run build // run webpack  server.bundle.js 

##### other programming langages
- python : pip

##### big picture
code ---publish-> artifactory ---pull(wget, curl)-> server

Nexus: 
```
$ docker volume create --name nexus-data
$ docker run -d -p 8081:8081 --name nexus -v nexus-data:/nexus-data sonatype/nexus3
$ cat /opt/sontype/sonatype-work/nexus3/admin.password to get the admin pwd

version: '3'
services:
  nexusServer:
    image: sonatype/nexus3
    ports:
     - 8081:8081
    volumes:
     - nexus-data1:/nexus-data
volumes: 
    nexus-data1
```



#### Docker
- Is an alternative for all other artifact types(docker image is an artifact) but I still need to build tha app
- no need to build and move different artifact type(jar, war, zip)
- just one artifact type -Docker image
- we build those Docker images from the app
- no need to install dependencies on the server (execute install command inside Docker image)
- no need to install npm or java on the server (everything within the image)

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


##### VM vs Container
- VM is a hardware abstruction: it takes physical CPUs and RAM from a host, and divides and shares it accross several sammer VMS.
OS and application running inside the VM but the virtualization software usually has no real knowledge of that.
- Container is an application abstraction


- Container running instance of an image, each container is isolated from the other container even when created from the same image.

```
$ docker commit CONTAINER_ID
$ docker image tag IMAGE_ID imageName

$ docker build -it hello:v0.1 .
$ docker image inspect --format "{{ json .RootFS.Layers }}" 

```
- Layer : A Docker image is build up from a series of layers. Each layer represetnts an instruction in the image's dockerfile.
- Dockerfile: file container all the commands to create an image
- Volumes: docker container layer that allows data to persist and be shared 

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

----------------------------------------------------------------------------------------------------------------------------------------

## Cloud & IaaS 

- Infrastructure as a service
- Move my physical infrastructure to cloud , delegate infrastructure management to the other cloud provider: Azure, AWS, digital ocean 

- Cloud computing is model for enabling ubiquitous, convenient, on-demand network access to a shared pool of configurable computing resources
  (network, servers, storage, applications and services) that can be rapidly provisioned and released with minimal management effort or service
provider interaction .
- This cloud model is composed of five essential characteristics, three service models and four deployment models.
- Essential characteristics:
1. on-demand self-service
2. Broad network access
3. Resource pooling
4. Rapid elasticity 
5. Measured service

- Service Model:
1. Saas
2. Paas
3. Iaas: the capability provided to the consumer is to provision processing, storage, network and other fundamental computing resources where consumer
   is able to deploy and run arbitrary software which include os and apps.


----------------------------------------------------------------------------------------------------------------------------------------
## Artifact Repository manager with nexus 
- Artifact: apps built into a single file (jar, war, zip, tar, etc)
- Artifact rep = storage of those artifacts
- Repository for each file/artifact type, one for jar, one for war , one for docker images ,extra

- Artifact Repository Manager (one app to mange all repository)
Nexus, Jfrog 
upload and store different built artifacts
retrieve(download) artifact later
central storage (shareable)
host own repositories
proxy repository 

- There is also public repository mangers like maven central repository , npm etc

- Nexus ARM: apt, helm, maven npm ,nuget, go RubyGems, PyPi, Docker
- Can integrate with LDAP
- Powerful Rest Api
jenkins --push---> Nexus <--pull--- server
- backup and restore
- metadata tagging (labelling and tagging artifacts)
- cleanup policies : automatically delete files that match certain condition 
- search functionality 


- Nexus repo type:
1. Proxy repo: check if the library exist in the nexus (already esit in the cache) or  grab to remote repo like maven and put it in the cache
2. group: combine multiple repository into one repo, use single url to access
3. hosted : releases , snapshots


----------------------------------------------------------------------------------------------------------------------------------------
### Docker

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


#### Main Docker commands
see docs
#### Develop with container
Dev ---> commit (git)---> Jenkins ---> Jenkins CI i produce artifact and build docker image from that artifact---->  push to Private Repo --> docker image deployed on pord server

#### Docker compose
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
#### Dockerfile
#### Push to private Docker repo
#### Deploy containerized application
#### Persist data in Docker - Volumes
#### Deploy and setup Nexus as Docker container
#### Push and fetch Docker images from Nexus repo



----------------------------------------------------------------------------------------------------------------------------------------
## Build Automation 
- Build app, build docker images , push to nexus 
dev --push-->Git repo --> should be tested and build
- publish new release: login to correct Docker repo, setup test environment to run tests, execute tests, build image and publish to repo

===>  Solution automate this process (detected server to automatise build process)

A detected server automatically will be trigger to run the build process , also will store all credentials: 
- Test code -> build app -> push to repo -> deploy to server
Tools: Jenkins

- What i can do with jenkins: Fucking middle man
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


```
http://localhost:8080/env-vars.html/
```

Jenkins File:
```
pipeline {
    agent any
	parameters{
		string(name: 'VERSION', defaultValue:'', description:'version to deoploy on prod')
		choice(name: 'VERSION', choices:['1.1.0','1.2.0', '1.3.0'], description: '')
	booleanParam(name:'executeTests', defaultValue: true, description: '')
	}
	tools{
		maven 'Maven'
		gradle
		jdk
	}
    environment{
        NEW_VERSION = '1.3.0'
        SERVER_CREDENTIALS = credentials('id_reference_of_credentials')
// define credentials in jenkisn UI , install plugin "Credentials Binding" No need to defined here in environment ,
// it's possible to call credentials in a specific stage: using withCredentials([
	usernamePassword(credentials: 'id_reference_of_credentials', usernameVariable: USER, passwordVariable: PWD)
])
    }


    stages {
        stage('build') {
            steps {
                echo 'Bild app !!!'
                echo "building version ${NEW_VERSION}"
 
            }
        }
         stage('test') {
		when {
					expression{
							params.executeTests == true
						}
            
            steps {
                echo 'Test app !!!'
            }
        }
             stage('deploy') {
		steps {
                echo 'Deploy app !!!'
            }
        }
    }
}


```
----------------------------------------------------------------------------------------------------------------------------------------

## AWS 
- amazon web services: collection of mini servercies.
- category:
1. compute:
	- EC2: virtual servers to rent
	- Lambda
	- 
2. storage:
	- S3
	- AWS backup

3. database
4. networking & content delivery
	- VPC
5. security, identity, & compliance
	- IAM
	- resource access manager
6. internt of things
7. containers
	- ECR
	- Elastic container service
	- Elastic kubernetes service

	
#### IAM identity and access management
- manage access to AWS services and resources
- who allowed to access
- create and manage AWS users and groups
- assign policies (set of permissions)

- ROOT user has unlimited privileges
- create an admin user with less privileges (create users assign resources ...)
- system users: application like jenkins 
- Other users 

- IAM Roles 
- Policies cannot be assigned to users directy to users, first it should assigned to role and  that role will be assigned to a user
- Role for each service. Policy specific to service 


/bin/bash: q: command not found
- Availability Zones: muliple datacenter in same region for replication ,if somthing happen to one datacenter in a region there is other to replace
  it.







----------------------------------------------------------------------------------------------------------------------------------------

## Kubernetes:
- Open source container orchestration tool 
- hepls manage containerized applications, in defferent envirement cloud physical , Virtual , hypred

Gananties: 
- High availabiity or no downtime
- Scalability  or hight performance
- Disaster recovery - backup and restore

#### K8s Components:

1. Node: virtual or physical server 
2. pod: smallest unit of k8s abstraction over container : layer on top the container [Pod[container]]
- One application per Pod (many container)
- Each Pod gets its own IP address (ensure communication between each other)
- Pods are ephemeral , each time a pod crashes if it restart then il we get an new ip address or re-creation


3. Service 
- permanent IP address  to pods
- life cycle of pod and services not connected
External service , Internal service(like db) 

4. Ingress ---> forward to External service

5. configMap and secret
ConfigMap : external configuration of my application like db url extera , credentials 
Secret like configMap stored in base64 encoded instead of plain text

6. Volumes
attach physical storage to the pod it can be storage on local machine , remote storage (cloud, on premise storage)

like external hard drive

7. Deployment:
- what happen when a pod crashes:  so we have to replicate everything on multiple nodes
- we need services(permanent IP , load balancer) and deployment (define a blueprints for pods  specifies the number of replicas also scale up or down number of pods)
- we can't replicate database in deployment 

8. Stateful set
- only for databases
- take care of replicating pods and scale them up and down
-  Avoid data intestacy  read and write are synchronized 

- Difficult to work with (another approach to host an external database)




