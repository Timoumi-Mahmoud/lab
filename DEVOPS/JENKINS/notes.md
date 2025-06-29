## Build Automation & CI/CD with Jenkins

What is build Automation ?
Deploy jenkins with docker compose
Install build tools in jenkins 
Create freestyle job (credentials, plugins, git)
Integreate Docker in jenkins: (Build image, publish to private repo)
Build scripted Pipeline
Jenkinsfile
Build multibranch Pipeline
Configure automated versioning
Jenkins Shared Library ( create jenkins shared library to make pipeline code reusable in other porjects)

[Dev]--> [Increment version] [Build app] [Build Image] [Push to Docker repo / Nexus] 

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
sets up jenkins cluster
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


