# CAML
Continuous Audit Metrics Library

## Basic elements of CAML source-distribution
## How is the repo organized?

* Continuous Audit Metrics Catalog
    * The catalog published by Cloud Security Alliance of metrics (may be linked from an external repository)
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
        * 
* codebase under src
    * commonlibs - 
    * main - 

## How to get started?
## How to setup?

1. clone the repo and all that
2. have docker installed
3. use docker compose to instantiate the cluster

    ```
    docker compose -f Docker/docker-compose.caml_dev.yaml up -d
    ```

    The docker-compose file runs both the YugabyteDB and the caml_dev images. 
    
    The caml_dev images is configured as specified in Dockerfile.caml_dev. This includes cloneing a copy of the repository into the image for in image development. What you need to do is **tell docker-compose which repository to use**. There are two options: 
        1. the **docker-compose command automatically used the value in the .env file** located in the same directory. This file is provided and specifies the main CAML repository on github. You can of course change this to any fork you are developing on. 
        2. You could specify the value on the command line when launching docker compose: 
        ```
        gitrepo=http://github.com/continube/CAML  docker compose -f Docker/docker-compose.caml_dev.yaml up -d
        ```
        Docker's environment/.env variable behavior is super overloaded. see https://docs.docker.com/compose/compose-file/compose-file-v3/#env_file as a starting point if you have trouble. 
 <br>       
    _3.5_ OPTIONAL: the compose file references the Docker/Dockerfile.caml_dev which you _could_ also build on its own.
    * Build a docker image from the root of the cloned repository
    ```
    docker build -f Docker/Dockerfile.caml_dev -t caml_dev_manual --build-arg gitrepo=https://github.com/continube/CAML.git .
    ```
    * You can explore directly with:
    ```
        docker run -it --rm caml_dev_manual
        cd CAML/src/main
        ./camldataservice 
    ```
    or:
    ```
        docker run -it --rm caml_dev_manual CAML/main/camldataservice
    ```
 
 5.  OPTIONAL: Do development in the container.
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


## References






## How to contribute?




