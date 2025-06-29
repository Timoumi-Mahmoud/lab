###
## What is an Artifact Repository: 
Artifacts are application built into a single file(war, jar, zip, tar) 
Articats Repository is where to store those artifacts
For each file/artifact type it should have a different repository

==> Artifact Repository Manager: One application (Nexus, JFrog) to mange all artifacts which allows upload and store different build artifacts and retrieve (download) artifacts later.

there is also public repository managers like mnv repository, npm, pip
Nexus support those formats: apt, composer, docker, Go, Helm, Maven, npm , NuGet pypi, R, RubyGems, yum...
Can be integrated with LDAP
offer a flexible and powerfull REST API for integration with other tools(part of build automation CI/CD )

[JENKINS] ---push---->[NEXUS] <---pull----[SERVER]

Backup and restore

mutli-format support (zip, tar,docker, etc)
metadata tagging(labelling and tagging artifacts) release, dev versions
Cleanup policies (every commit on master branch create an artifact so we need to clean up keep the latest one and delete those with a condition like date)




## Installation
```bash
$ docker volume create --name nexus-data
$ docker run -d -p 8081:8081 --name nexus -v nexus-data:/nexus-data sonatype/nexus3
$ cat /opt/sontype/sonatype-work/nexus3/admin.password to get the admin pwd
```
Two directories : 
Nexus folder contains runtime and application of nexus (when upgrade to a newer version this folder wil be changes)
Sonyatype-work contains config for Nexus and data (ip address, logs ,uploaded files and metadata, config) can be used as backup 
```yaml
version: '3'
services:
  nexusServer:
    image: sonatype/nexus3
    ports:
     - 8081:8081
    volumes:
     - nexus-data1:/nexus-data
volumes: 
    nexus-data1:
```

## Repository Types

Three types: hosted, group, proxy
1. Proxy: a repository linked to a remote Repository like mvn (link to remote repository) for example if I try to download a library it will directly go to the proxy repo instead of centeral mvn and check if the library already exists in nexus if use then my application will be retrive it from Nexus if No then the request will be forwared to the remote repoistory and then it will be stored in nexus for future requests (act as cache)
=> save network bandwidth and time 
(maven-central, nuget.org-proxy, docker proxy)

2. Hosted: Primary storage for artifacts and components 
maven-snapshots, maven-releases, nuget-hosted

3. Group: combine multiple repository and access them via single url

## Upload Jar (Spring+ mvn)

1. put settings.xml  under ~/.m2/ which contains username, password of nexus user & the id of spring project
2. update pom.xml (add maven plugin and nexus url + id)
3. mvn package , mvn deploy

## Nexus API

Query Nexus Reposiotyr for different information 
Which components ? what are the verson ? which repository ? ...
This information is needed in my CI/CD Pipeline
(curl or wget)
example: 

```bash
$ curl -u mahmoud:0000@Lambda -X GET 'http://localhost:8081/service/rest/v1/repositories'

$ curl -u mahmoud:0000@Lambda GET 'http://localhost:8081/service/rest/v1/components?repository=maven-snapshots' 

$ curl -u mahmoud:0000@Lambda GET 'http://localhost:8081/service/rest/v1/components/<id-component>'
```

## Blob Storage
Internal storage for binary files
It can be local storage(file type) or cloud storage like s3
[type, state(started|failed), blob counts(vol-n),Total size, Available space ] 

Blob store can't be modified
Blob store used by a repository can't be deleted (first delete the repository)

## Components vs Assets

com:
    example:
        java-maven-app: Component 
        xxx.jar   assets
        .....

        myapp-maven-app: Component 

- Components: abstract of what I am uploading  
- Assets: actual physical packages/files 
1 component=1 or more assets

! in Docker : Docker format gives assets unique identifiers(docker layer == Assets) 
Example: 2 Docker images=> 2 components, but share the same assets(layer)


## Cleanup
Cleanup policies allow you to automatically delete unused components and free up storage space.
I can schedule cleanup policies to run at specific intervals and define cleanup criteria for selecting components to delete.

Release Type
Remove components that are of the following release type: releases | pre-release | snapshots

Component Age
Components published over “x” days ago 

Component Usage
Components downloaded in “x” amount of days 

Asset Name Matcher
Remove components that have at least one asset name matching the following regular expression pattern:

1. create a cleanup policies 
2. attach cleanup policies to a repository
3. it will be scheduled (see Tasks)
4. to actually delete them not only soft delete (Task - Admin compact blob store)
