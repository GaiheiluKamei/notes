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

