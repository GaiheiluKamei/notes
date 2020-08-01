# Docker Compose

## Demystifying declarative versus imperative

> **Imperative**: This is a way in which we can solve problems by specifying the exact procedure that has to be followed by the system.  
> **Declarative**: This is a way in which we can solve problems without requiring the programmer to specify an exact procedure to be followed.  

## Running a multi-service app

Let's have a look at the content of a simple `docker-compose.yml` file:  

```yml
version: "2.4"
services:
 web:
    image: fundamentalsofdocker/ch11-web:2.0
    build: web
    ports:
    - 80:3000
 db:
    image: fundamentalsofdocker/ch11-db:2.0
    build: db
    volumes:
    - pets-data:/var/lib/postgresql/data

volumes:
 pets-data:
```  

The lines in the file are explained as follows:  

- `version`: In this line, we specify the version of the Docker Compose format we want to use.  
- `services`: In this section, we specify the services that make up our application in the `services` block. In our sample, we have two application services and we call them `web` and `db`.  
- `web`: The `web` service is using an image called `fundamentalsofdocker/ch11-web:2.0`, which, if not already in the image cache, is built from the `Dockerfile` found in the `web` folder. The service is also publishing container port `3000` to the host port `80`.  
- `db`: We are mounting a volume called `pets-data` into the container of the `db` service.  
- `volumes`: The volumes used by any of the services have to be declared in this section. In our sample, this is the last section of the file. The first time the application is run, a volume called `pets-data` will be created by Docker and then, in subsequent runs, if the volume is still there, it will be reused.  

> Note that we are using version 2.x of the Docker Compose file syntax. This is the one targeted toward deployments on a single Docker host. There exists also a version 3.x of the Docker Compose file syntax. This version is used when you want to define an application that is targeted either at Docker Swarm or Kubernetes.  

**Building images with Docker Compose**:  

> `docker-compose build`  

If we enter the preceding command, then the tool will assume that there must be a file in the current directory called `docker-compose.yml` and it will use that one to run.  

**Running an application with Docker Compose**:  

> `docker-compose up`  

We can also run the application in the background. All containers will run as daemons.  

> `docker-compose up -d`  

To list all services that are part of the application:  

> `docker-compose ps`  

To stop and clean up the application, we use the `docker-compose down` command:  

> `docker-compose down`  

If we also want to remove the volume for the database:  

> `docker volume rm ch11_pets-data`  

Alternatively, instead of using the two commands, we can combine them into a single command:  

> `docker-compose down -v`  

Why is there a `ch11` prefix in the name of the volume? Docker Compose prefixed all names with the name of the parent folder of the `docker-compose.yml` file plus an underscore. If you don't like this approach, you can define a project name explicitly, for example:  

> `docker-compose -p my-app up`  

which uses a project name my-app for the application to run under.  

## Scaling a service

Running more instance is also called **scaling up**. We can use this tool to scale our `web` service up to, say, three instances:  

`docker-compose up --scale web=3`  

If we do this, we are in for a surprise:  

> `ERROR: for ch11_web_3 ... Bind for 0.0.0.0:80 failed: port is already allocated`  

We can just let Docker decide which host port to use for each instance. If, in the `ports` section of the `compose` file, we only specify the container port and leave out the host port, the Docker automatically selects an ephemeral port.  

```yml
ports:
  - 3000
```  

To push all images to Docker Hub, we can use `docker-compose push`. We need to be logged in to Docker Hub so that this succeeds.  

## Using Docker Compose overrides

To this basic `docker-compose.yml` file:  

```yml
version: "2.4"
services:
  web:
    image: fundamentalsofdocker/ch11-web:2.0
  db:
    image: fundamentalsofdocker/ch11-db:2.0
    volumes:
      - pets-data:/var/lib/postgresql/data

volumes:
  pets-data:
```  

If we want to override some environment variables, we can create a new `docker-compose-ci.yml`:  

```yml
version: "2.4"
services:
  web:
    build: web
    ports:
      - 5000:3000
    environment:
      POSTGRES_PASSWORD: ci-pass
  db:
    build: db
    environment:
      POSTGRES_PASSWORD: ci-pass
```  

Then we can run this application with the following command:  

> `docker-compose -f docker-compose.yml -f docker-compose-ci.yml -d --build`  

When using environment variables, note the following precedence:  

- Declaring them in the Docker file defines a default value  
- Declaring the same variable in the Docker Compose file overrides the value from the Dockerfile  

> Had we followed the standard naming convention and called the base file just `docker-compose.yml` and the override file `docker-compose.override.yml` instead, we could have started the application with `docker-compose up -d` without explicitly naming the compose files.  

## Summary

Typically, developers and CI servers work with single hosts and those two are the main users of Docker Compose.
