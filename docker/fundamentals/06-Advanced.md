# Advanced Docker Usage Scenarios

## All of the tips and tricks of a Docker pro

### Keeping your Docker environment clean

First, we want to learn how we can delete dangling images. According to Docker, dangling images are layers that have no relationship to any tagged images. Such image layers are certainly useless to us and can quickly fill up our disk-it's better to remove them from time to time. Here is the command:  

> `docker image prune -f`  

Stopped containers can waste precious resources too, the following command will remove all the stopped containers:  

> `docker container prune --force`  

### Running Docker in Docker

At times, we may want to run a container hosting an application that automates certain Docker tasks. How can we do that? The Docker Engine and The Docker CLI are installed on the host, ye the application runs inside the container. Well, from early on, Docker has provided a means to bind-mount Linux sockets from the host into the container. On Linux, sockets are used as very efficient data communications endpoints between processes that run on the same host. The Docker CLI uses a socket to communicate with the Docker Engine; it is often called the Docker socket. If we can give access to the Docker socket to an application running inside a container then we can just install the Docker CLI inside this container, and we will then be able to run an application in the same container that uses this locally installed Docker CLI to automate container-specific tasks.  

> **Note**: It is important to note that here we are not talking about running the Docker Engine inside the container but rather only the Docker CLI and bind-mount the Docker socket from the host into the container so that the CLI can communicate with the Docker Engine running on the host computer.  

Assume that we have the following script, called `pipeline.sh`:  

```sh
#! /bin/bash
# *** Sample script to build, test and push containerized Node.js applications ***
# build the Docker image
docker image build -t $HUB_USER/$REPOSITORY:$TAG .
# Run all unit tests
docker container run $HUB_USER/$REPOSITORY:$TAG npm test
# Login to Docker Hub
docker login -u $HUB_USER -p $HUB_PWD
# Push the image to Docker Hub
docker image push $HUB_USER/$REPOSITORY:$TAG
```  

We want to run that script inside a builder container.Since the script uses the Docker CLI, our builder container must have the Docker CLI installed, and to access the Docker Engine, the builder must have the Docker socket bin-mounted. Let's start creating a Docker image for such a builder container:  

```sh
# create a `builder` folder
mkdir builder && cd builder

# inside this folder, create a Dockerfile
FROM alpine:latest
RUN apk update && apk add docker
WORKDIR /usr/src/app
COPY . .
CMD ./pipeline.sh  

# make the file an executable
chmod +x ./pipeline.sh

# build an image
docker image build -t builder .

# we can use the Node.js before we defined, run a container
docker container run --rm --name builder \
-v /var/run/docker.sock:/var/run/docker.sock \
-v "$PWD":/usr/src/app \
-e HUB_USER=<user> \
-e HUB_PWD=<password>@j \
-e REPOSITORY=ch08-sample-app \
-e TAG=1.0 builder
```  

Notice how, in the preceding command, we mounted the Docker socket into the container with the `-v /var/run/docker.sock:/var/run/docker.sock`. This is only one of the many use cases where it is very useful to be able to bind-mount the Docker socket.  

### Formatting the output of common Docker commands

> `docker container ps -a --format "table {{.Names}}\t{{.Image}}\t{{.Status}}`  

### Filtering the output of common Docker commands

The format of filters is straightforward and of the type `--filter <key>=<value>`. If we need to combine more than one filter, we can just combine multiple of these statements.  

> `docker image ls --filter dangling=false --filter "reference=*/*/*/:latest"`  

