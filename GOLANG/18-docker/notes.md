## Run Go in Docker
- run golang porject inside docker container, and mount the current directory as a volume to the container.

```bash 
docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
docker build --target runtime . -t runtime-app
docker build --target build . -t customer-app

docker run runtime-app
```


* Packages:
    - source files in the same directory that are compiled together.
    - have visibility on all source files in the same package.
* Modules:
    - collection of packages that are released together.

* go.mod: declares module path + import path for packages (where to download them)=> where it tracks dependency.

* go install: this command builds the app, producing an executable binary.it then installs that binary as $HOME/go/bin/app
```bash
go install github.com/Timoumi-Mahmoud/GOLANG/18-docker/app
```



