# Debugging Code Running in Containers

We're going to introduce techniques commonly used to allow a developer to evolve, modify, debug, and test their code while running in a container. With these techniques at hand, you will enjoy a frictionless development process for applications running in a container, similar to what you experience when developing applications that run natively.  

## Evolving and testing code running in a container

We can achieve a significant reduction in the friction in the development process by mapping the source code in the running container (with `--volume` parameter). We can now add new or modify existing code and test it without having to build the container image first. Yet, there is still a bit of friction left in play. We have to manually restart the container every time we want to test some new or modified code.  

## Auto restarting code upon changes

If you have been coding for a while, you will certainly have heard about helpful tools that can run your applications and restart them automatically whenever they discover a change in the code base. For `Node.js` applications, the most popular such tool is `ndoemon`, we can build a image and start a container with the following steps:  

```sh
# use npm create a Node.js project, accept all the defaults while creating.
mkdir -p ~/fod/node-test && cd ~/fod/node-test && npm init

# we want to use the Express.js library in our Node application, install it:
npm install express --save

# start VS Code from within this folder:
code .

# in VS Code, create a new `index.js` file and add this code snippet to it.
const express = require('express');
const app = express();

app.listen(3000, '0.0.0.0', ()=>{
    console.log('Application listening at 0.0.0.0:3000');
})

app.get('/', (req,res)=>{
    res.send('Sample Application: Hello World!');
})

# create a Dockerfile with these contents:
FROM node:latest
RUN npm install -g nodemon
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
CMD nodemon

# build the image:
docker image build -t sample-app .

# run a container:
docker container run --rm -it -v $(pwd):/app -p 3000:3000 sample-app-dev
```  

Now, while the application is running in the container, change the `index.js` output content, and test with `curl localhost:3000` command, the application inside the container is automatically restarted.  

## Further reading  

> [Live debugging with Docker](https://www.docker.com/blog/live-debugging-docker/)  
> [Debug apps in a local Docker container](https://docs.microsoft.com/en-us/visualstudio/containers/edit-and-refresh?view=vs-2019)  
> [Debug your java applications in Docker using IntelliJ IDEA](https://blog.jetbrains.com/idea/2019/04/debug-your-java-applications-in-docker-using-intellij-idea/)