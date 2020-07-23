# Single-Host Networking

## Working with the bridge network

When the Docker daemon runs for the first time, it creates a Linux bridge and calls it `docker0`. This is the default behavior and can be changed by changing the configuration. Docker then creates a network  with this Linux bridge and calls the network `bridge`.  

We can list all the networks on the host with the following command:  

> `docker network ls`  

Let's look a little bit deeper into what this `bridge` network is all about:  

> `docker network inspect bridge`  

```json
# part of output generated when inspecting the docker bridge network
{
    "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        }
}

```

The output is nothing new. But let's have a look at the **IP address management(IPAM)** block. IPAM is a piece of software that is used to track IP addresses that are used on a computer. The important part of the `IPAM` block is the `Config` node with its values for `Subnet` and `Gateway`. The subnet for the bridge network is defined by default as `172.17.0.0/16`. This means that all containers attached to this network will get an IP address assigned by Docker that is taken from the given range, which is `172.17.0.2` to `172.17.255.255`. The `172.17.0.1` address is reserved for the router of this network whose role in this type of network is taken by the Linux bridge.  

We are not limited to just the bridge network, as Docker allows us to define our own custom bridge networks. This is not just a feature that is nice to have, but it is a recommended best practice to not run all containers on the same network. Instead, we should use additional bridge networks to further isolate containers that have no need to communicate with each other. To create a custom bridge network called `sample-net`, use the following command:  

> `docker network create --driver bridge sample-net`  

If, for some reason, we want to specify our own subnet range when creating a network, we can do by using the `--subnet` parameter:  

> `docker network create --driver bridge --subnet "10.1.0.0/16" test-net`  

Create a container attach it to a network:  

> `docker container run --name c3 -d --network test-net alpine ping 127.0.0.1`  

We learned that a container can be attached to multiple networks:  

> `docker container run --name c3 -d --network test-net --network sample-net alpine 127.0.0.1`  

## The host and null network

When running business software in containers, there is no good reason to ever run the respective containers attached to the host's network. For security reasons, it is strongly recommended that you do not run any such container attached to the `host` network on a production or production-like environment.  

Sometimes, we need to run a few application services or jobs that do not need any network connection at all to execute the task at hand. It is strongly advised that you run those applications in a container that is attached to the `none` network. This container will be completely isolated, and is thus safe from any outside access.  

> `docker container run --rm -it --network none alpine:latest /bin/sh`  

## Running in an existing network namespace

Docker provides an additional way for us to define the network namespace that a container runs in. When creating a new container, we can specify that it should be attached to (or maybe we should say included) in the network namespace of an existing container. With this technique, we can run multiple containers in a single network namespace.  

This is useful when we want to debug the network of an existing container without running additional processes inside that container. We can just attach a special utility container to the network namespace of the container to inspect.  

```docker
# first, we create a new bridge network:
docker network create --driver bridge test-net

# next, we run a container attached to this network:
docker container run --name web -d \
--network test-net nginx:alpine

# finally, we run another container and attach it to the network 
# of our `web` container
docker container run -it --rm --network container:web \
alpine:latest /bin/sh

# since the new container is in the same network namespace as the web container
# running nginx, we're now able to access nginx on localhost!
# We can prove this by uding the `wget` tool
/ # wget -qO - localhost

# to clean up the container and network, we can use the following command
docker container rm --force web
docker network rm test-net
```  

Please also note that there is an important difference between running two containers attached to the same network and two containers running in the same network namespace. In both cases, the containers can freely communicate with each other, but **in the latter case, the communication happens over localhost**.  

## Mapping container ports

Let's see how we can actually map a container port to a host port. This is done when creating a container.  

```docker
# we can automatically done with the `-P` parameter:
docker container run --name web -P -d nginx:alpine

# find out which host port Docker is using by the following command:
docker container port web

# an alternative way:
docker container inspect web | grep HostPort

# the last way:
docker container ls

# sometimes we want to map a container port to a very specific host port
# we can do this by using the `-p` (--publish) parameter
docker container run --name web2 -p 8080:80 -d nginx:alpine
```

When using the UDP protocol for communication over a certain port, the `publish` parameter will look like `-P 3000:4321/udp`. Note that if we want to allow communication with both TCP and UDP protocols over the same port, then we have to map each protocol separately.  

## HTTP-level routing using a reverse proxy

### Using Traefik to reroute traffic

[Traefik](https://docs.traefik.io/) is a cloud-native edge router and it is open source, which is great for our specific case. It can be combined with Docker in a very straightforward way.  

Here is an simplified example:

```docker
docker container run --rm -d \
--name catalog \
--label traefik.enable=true \
--label traefik.port=3000 \
--label traefik.priority=10 \
--label traefik.http.routers.catalog.rule="Host(\"acme.com\") && \
PathPrefix(\"/catalog\")" acme/catalog:1.0
```

Let's quickly look at the four labels we define:  

- `traefik.enable=true`: This tells Traefik that this particular container should be included in the routing (default is `false`).
- `traefik.port=3000`: The router should forward the call to port `3000`.
- `traefik.priority=10`: Give this route high priority.
- the last label: The route must include the hostname and the path must start with `/catelog` in order to be rerouted to this services.  

> Please note the special form of the fourth label, Its general form is `traefik.http.routers.<service name>.rule`  
