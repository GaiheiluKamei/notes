# Orchestrators

## The tasks of an orchestrator

There are two quite different types of services that we might want to run in a cluster that is managed by an orchestrator. They are *replicated* and *global* services. In a replicated service, we will always be guaranteed to find the exact desired number of instances, while for a global service, we can be assured that on every worker node, there will always run exactly one instance of the service.  

In a cluster that is managed by an orchestrator, we typically have two types of nodes, *managers* and *workers*. A manager node is usually exclusively used by the orchestrator to manage the cluster and does not run any other workload. Work nodes, in turn, run the actual applications.  

> In Kubernetes, a global service is also called **DaemonSet**.  

Services meed to not just be able to scale up, but also to scale down when the workload goes down.  

