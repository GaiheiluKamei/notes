# Building and Deploying Go Code

Go's code is compiled into binaries, which allows for the modular deployment of new Go code as we iterate through code development. We can push this out to one or multiple places in order to test against different environments. Doing this will allow us to optimize our code to fully utilize the throughput that will be available to us in our system.  

We'll look at how the Go compiler builds binaries, and we'll use this knowledge to build right-sized, optimized binaries for the platform at hand. We will cover the following topics:  

> Building Go binaries
> Using `go clean` to remove object files  
> Using `go get` to download and install dependencies  
> Using `go mod` for dependency management
> Using `go list` to list packages and modules
> Using `go run` to execute programs
> using `go install` to install packages  

## Building Go binaries

Go's build system has quite a few options that can help the system operator add additional parameterization to their build strategy.  

In this chapter, we are going to investigate these different pieces of the Go build system. As we learn more about how these programs interoperate with one another, we will be able to see how to use them to our advantage to build slim, feature-filled binaries that will work as we expect them to on the supported architectures and operating systems.  

## Go build —— build your Go code

The invocation stanza for go build is as follows:  

> `go build [-o output] [build flags] [packages]`  

Packages can be defined as a list of go source files or they can be omitted. If a list of go source files is specified, the build program will use the list of files that were passed as a group that specifies a single package. If no packages are defined, the build program will validate that the packages with the directory can be built, but it will discard the result of the build.  

