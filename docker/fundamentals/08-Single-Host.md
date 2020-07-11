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