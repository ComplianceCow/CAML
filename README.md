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

3. the dev/runtime environment is defined by /docker/dockerfile

    Build a docker image from the root of the cloned repository
    ```
    docker build -f Docker/Dockerfile.caml_dev -t caml_dev .
    ```
    The dockerfile does an **ADD** to insert a **copy of the source** and builds an image you can run it and look around. The source that has been ADDed is in 'added.caml' so:
    ```
        docker run -it --rm caml_dev
        cd added.caml/main
        ./camldataservice 
    ```
    or just:
    ```
        docker run -it --rm caml_dev added.caml/main/camldataservice
    ```
    You could do some development right now. But the root image doesn't have editors and etc. So you'd be doing dev in a less than ideal environment. Also the supplimental containers for the database and etc aren't running yet. Lets keep going: 

4. use docker compose to instantiate the cluster

    ```
    docker compose -f Docker/docker-compose.caml_dev.yaml up
    ```

    The docker-compose file runs both the YugabyteDB and the caml_dev images. When running caml_dev it mounts "../src/:/go/src/volume.caml" (relative path from the docker-compose.yaml file). This results in /go/src/volume.caml existing in the container. This is a mount of the current source directory structure. 
    
    Now you can use any local editors as you wish. 

    But you still need to build and run within the limited environment of the caml_dev container. To do real development you ...
    * TODO option 0: do the build steps shown in Dockerfile.caml_dev to verify build works
        ```
        cd /go/src/volume.caml/main
        go mod download
        go mod tidy -compat=1.17 && go get -d -v
        CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o camldataservice .
        ```
    * TODO option 1: add to the dockerfile necessary commands to install a debug environment w/in the container etc
    * TODO option 2: use a toolchain like vscode. you can attach to the container and do development that way. vscode can handle most of this for you: in the vscode UI:
        * "attach to running container" (select caml_dev)
        * "install all" at the vscode prompts for go-outline and gopls 
        * re-evaluate your security brain: did you really just install all from some random vscode pointed respositories? 
        * anyway, now you can run and debug main.go w/ breakpoints and stuff 
    * TODO whatever else you want


## References






## How to contribute?




