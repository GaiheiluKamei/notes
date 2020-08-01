# Zero-Downtime Deployments and Secrets

## Zero-downtime deployment

### Rolling updates

When we declare a service in a stack file, we can define multiple options that are relevant in this context. Let's look at a snippet of a typical stack file:  

```yaml
ersion: "3.5"
services:
 web:
   image: nginx:alpine
   deploy:
     replicas: 10
     update_config:
       parallelism: 2
       delay: 10s
...
```  

In this snippet, `parallelism` defines the batch size of how many replicas are going to be updated at a time during a rolling update. `delay` defines how long Docker Swarm is going to wait between updating individual batches.  

If we just want to deploy a single service, this might be the preferred way of doing things, Let's look at such a `create` command:  

> `docker service create --name web --replicas 10 --update-parallelism 2 --update-delay 10s nginx:alpine`  

This command defines the same desired state as the preceding stack file.  

### Health checks

We can define a health check for a service in the `Dockerfile` of its image:  

```dockerfile
FROM alpine:3.6
...
HEALTHCHECK --interval=30s \
    --timeout=10s
    --retries=3
    --start-period=60s
    CMD curl -f http://localhost:3000/health || exit 1
...
```  

We can also define the health check in the stack file that we use to deploy our application into Docker Swarm:  

```yaml
version: "3.5"
services:
  web:
    image: example/web:1.0
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 60s
...
```  

If there are health check-related settings defined in the image, then the ones defined in the stack file override the ones from the `Dockerfile`.  

### Rollback

A rollback can be looked at as a reverse update. Here, we have the stack file that we used previously, but this time with some rollback-relevant attributes:  

```yaml
version: "3.5"
services:
  web:
    image: nginx:1.12-alpine
    ports:
      - 80:80
    deploy:
      replicas: 10
      update_config:
        parallelism: 2
        delay: 10s

        failure_action: rollback
        monitor: 10s

    healthcheck:
      test: ["CMD", "wget", "-qO", "-", "http://localhost"]
      interval: 2s
      timeout: 2s
      retries: 3
      start_period: 2s
```  

Now, under deploy, we have a new entry called `monitor`. This entry defines how long newly deployed tasks should be monitored for health and whether or not to continue with the next batch in the rolling update. We also have a new entry, `failure-action`, which defines what the orchestrator will do if it encounters a failure during the rolling update, such as that the service is unhealthy. By default, the action is just to stop the whole update process and leave the system in an intermediate state.  

## Storing configuration data in the swarm

```sh
# first, we start a simple string value
$ echo "Hello World" | docker config create hello-config -

# Let's see what we got
$ docker config ls

# we can do more and even output the content of a config
$ docker config inspect hello-config

# our string encoded as `base64`
$ echo 'SGVsbG8gd29ybGQK' | base64 -d
```  

Now, Let's define a somewhat more complicated Docker config.  

```sh
# Let's create a file called `my-app.properties`:  
username=pguser
database=products
port=5432
dbhost=postgres.acme.com

# create a Docker config called `app.properties`:
docker config create app.properties ./my-app.properties

# we can use this command to get the clear text value of the config:
docker config inspect app.properties | jq .[].Spec.Data | xargs echo | base64 -d

# let's create a Docker service that uses the preceding config
docker service create --name nginx \
  --config source=app.properties,target=/etc/my-app/conf/app.properties,mode=0440 \
  nginx:1.13-alpine
```  

## Protecting sensitive data with Docker secrets

By default, secrets are mounted into the container at `/run/secrets`.  

### Creating secrets

> `echo "sample secret value" | docker secret create sample-secret -`  

Please note the hyphen at the end of the `docker secret create` command. This means that Docker expects the value of the secret from standard input.  

Alternatively, we can use a file as the source for the secret value:  

> `docker secret create other-secret ~/my-secrets/secret-value.txt`  

Once a secret has been created, there is no way to access the value of it. We can list all our secrets:  

> `docker secret ls`  

### Using a secret

```docker
docker service create --name web -p 8000:8000 \
  --secret source=api-secret-key,target=/run/my-secrets/api-secret-key \
  fundamentalsofdocker/whoami:latest
```  
