# Debugging Code Running in Containers

We're going to introduce techniques commonly used to allow a developer to evolve, modify, debug, and test their code while running in a container. With these techniques at hand, you will enjoy a frictionless development process for applications running in a container, similar to what you experience when developing applications that run natively.  

## Evolving and testing code running in a container

We can achieve a significant reduction in the friction in the development process by mapping the source code in the running container (with `--volume` parameter). We can now add new or modify existing code and test it without having to build the container image first. Yet, there is still a bit of friction left in play. We have to manually restart the container every time we want to test some new or modified code.