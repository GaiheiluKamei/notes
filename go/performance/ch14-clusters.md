# CLusters and Job Queues

Clustering and job queues in Go are good ways to get distributed systems to work synchronously and deliver a consistent message. Distributed computing is difficult and it becomes very important to watch for potential performance optimizations within both clustering and job queues.  

Learning about different clustering systems can help you identify large groups of data and how to accurately classify them in your datasets. Learning about queueing systems will help you move large amounts of information from your data structures into specific queueing mechanisms in order to pass large amounts of data to different systems in real time.  

## Clustering in Go

Clustering is a methodology that you can use in order to search for consistent groups of data within a given dataset. Using comparison techniques, we can look for groups of items within dataset that contain similar characteristics. These individual datapoints are then divided into clusters. Clustering is commonly used in order to solve multi-objective problems.

There are two general classifications of clustering, both of which have distinct subclassifications:  

> - **Hard clustering**: The datapoints within the dataset are either explicitly a part of a cluster or not explicitly part of a cluster. Hard clustering can be further classified as follows:  
>   - **Strict partitioning**: An object can belong to exactly one cluster.  
>   - **Strict partitioning with outliers**: Strict partitioning, which also includes a concept that objects can be classified as outliers (meaning they belong to no cluster).  
>   - **Overlapping clustering**: Individual objects can be associated with one or more clusters.  
> 
> - **Soft clustering**: Datapoints are assigned to probability that they are associated with a particular cluster based on explicit criteria. They can be further classified as follows:  
>   - **Subspace**: Clusters use a two-dimensional subspace in order to be further classified into two dimensions.
>   - **Hierarchical**: Clustering using a hierarchical model; and object that is associated with a child cluster is also associated with the parent clusters.  

There are also many different algorithm types that are used for clustering, SOme examples are shown in the following table:  

|**Name**|**Definition**|  
|---|--|
|Hierarchical|Used to attempt to build a hierarchy of clusters. Usually based on a top-down or a bottom-up approach, attempting to segment datapoints either from one to many clusters (top-down) or many to few clusters (bottom-up).  
|Centroid| Used to find a specific point location that acts as the center of a cluster.  
|Density| Used to look for places in the dataset that have dense regions of datapoints.
|Distribution| Used to utilize distribution models to order and classify datapoints within a cluster.
  
In this book, we're going to focus on hierarchical and centroid algorithms as they are commonly used in computer science (namely in machine learning).  

### K-nearest neighbors

Hierarchical clustering is a clustering method in which ab object that is associated with a child cluster is also associated with the parent clusters. The algorithm begins with all of the individual datapoints in the data struct being assigned to individual clusters. The nearest clusters to one another merge. This pattern continues until all the datapoints have an association with another datapoint. Hierarchical