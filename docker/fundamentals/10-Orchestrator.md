# Orchestrators

## The tasks of an orchestrator

There are two quite different types of services that we might want to run in a cluster that is managed by an orchestrator. They are *replicated* and *global* services. In a replicated service, we will always be guaranteed to find the exact desired number of instances, while for a global service, we can be assured that on every worker node, there will always run exactly one instance of the service.  

In a cluster that is managed by an orchestrator, we typically have two types of nodes, *managers* and *workers*. A manager node is usually exclusively used by the orchestrator to manage the cluster and does not run any other workload. Work nodes, in turn, run the actual applications.  

> In Kubernetes, a global service is also called **DaemonSet**.  

Services meed to not just be able to scale up, but also to scale down when the workload goes down.  

All these activities, where the orchestrator monitors the current state and automatically repairs the damage or reconciles the desired state, lead to a so-called **self-healing* system.  

The orchestrator needs to be able to update individual application services, batch-wise. This is also called **rolling updates**.  

The communication that happens in a cluster can be separated into three types. You talk about communication planes-management, control, and data planes:  

> The management plane is used by the cluster managers, or masters, to, for example, schedule service instances, execute health checks, or create and modify any other resources in the cluster, such as data volumes, secrets, or networks.  
>  
> The control plane is used to exchange important state information between all nodes of the cluster. This kind of information is, for example, used to update the local IP tables on clusters, which are used for routing purposes.  
>  
> The data plane is where the actual application services communicate with each other and exchange data.  

Normally, orchestrators mainly care about securing the management and control plane. Securing the data plane is left to the user, although the orchestrator may facilitate this task.  
