# CAML
Continuous Audit Metrics Library

## How to get started?

### Clone the repo and all that
### Have docker installed
### Instantiate the cluster:

Using docker compose:

```
docker compose -p caml_dev -f Docker/docker-compose.caml_dev.yaml up -d
```

The docker-compose file runs both the YugabyteDB and the caml_dev images. 
    
The caml_dev images is configured as specified in Dockerfile.caml_dev. This includes to clone a copy of the repository into the image for in image development. What you need to do is **tell docker-compose which repository to use**. There are two options:
    
   - The **docker-compose command automatically used the value in the .env file** located in the same directory. This file is provided and specifies the main CAML repository on github. You can of course change this to any fork you are developing on. 
   - You could specify the value on the command line when launching docker compose to quickly indicate your forked respository: 
       
```
GITREPO=http://github.com/yourusername/CAML  docker compose -f Docker/docker-compose.caml_dev.yaml up -d
```
        
(NOTE: Docker's environment/.env variable behavior is super overloaded. see https://docs.docker.com/compose/compose-file/compose-file-v3/#env_file as a starting point if you have trouble. An environment variable specified on the command line like this should overload whats in the .env file)        

## Explore!

To help get things going an instance of prometheus and grafana are installed. The way this works is:

* that caml_dev does its best to collect (or generate test data) and compute some example metrics
  * the test data is silly. it creates metrics like every second or two. very silly. but visible. 
* the example metrics are placed where node_exporter can make them available to prometheus
* prometheus stores the values in a time series database. The UI page for prometheus offers some very simple graph methods and ability to explore. 
* grafana provides a much more complete dashboarding. A dashboard for current metrics is automatically provisioned. 

### Join us! Develop a cool feature
* **Seems best**: use a toolchain like vscode. you can attach to the container and do development that way. vscode can handle most of this for you: in the vscode UI:
  * "attach to running container" (select caml_dev)
  * "install all" at the vscode prompts for go-outline and gopls 
  * re-evaluate your security brain: did you really just install all from some random vscode pointed respositories? 
  * anyway, now you can run and debug main.go w/ breakpoints and stuff 
  * because this is a cloned git repository vscode can handle all the git operations to commit changes. The first time you try this it will walk you through getting authorization credentials for vscode. 
* **Simplest test**: Do the build steps shown in Dockerfile.caml_dev to verify build work

```    
cd /go/src/volume.caml/main
go mod download
go mod tidy -compat=1.17 && go get -d -v
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o camldataservice .
```
* TODO: whatever else you want. Like you could, add to the dockerfile necessary commands to install a debug environment w/in the container etc. Basically do it all manually.

#### OPTIONAL (not recommended): instantiate parts of the cluster manually: 

The compose file references the Docker/Dockerfile.caml_dev which you _could_ also build on its own.
Build a docker image from the root of the cloned repository:

```
docker build -f Docker/Dockerfile.caml_dev -t caml_dev_manual --build-arg gitrepo=https://github.com/continube/CAML.git .
```
      
You can explore directly with:

```
docker run -it --rm caml_dev_manual
cd CAML/src/main
./camldataservice 
```
Or explore directly with:

```
docker run -it --rm caml_dev_manual CAML/main/camldataservice
```
 
## Basic elements of CAML repo

* Continuous Audit Metrics Catalog
    * The catalog published by Cloud Security Alliance of metrics 
    * TODO: (may be linked from an external repository)
* Inject processing to convert the catalog into CAML specific structures.       
    * The intention is to make this is lite as possible. 
    * see etc/appscript.js
    * controls.json
    * metrics.json
* YugabyteDB SQL database. measurement storage.
    * yb-standalone
    * used for storing measurements and computed metrics
* (discussed) evidence storage
    * MinIO object store service
* publishing service
    * regularly pushes the metrics to a CSA star
    * also has a 'file' sync that is regularly updated for testing and local use
    * (discussed) multiple publishing methods
     * e.g. to tableau or smartsheets or something that is appropriate for internal use 
* development environment
    * TODO: 
    * a runtime environment with the codebase loaded in
    * either by copying from codebase via Dockerfile
    * or by mounting the codebase 
* runtime environment
    * implemented by building the copied/mounted codebase
    * automatically launching:
     * API server 
* codebase under src
    * commonlibs - 
    * main - 



## References






## How to contribute?




