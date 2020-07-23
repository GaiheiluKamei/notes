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

