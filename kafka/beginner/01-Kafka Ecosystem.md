# Kafka Ecosystem

## What is Apache Kafka

Apache Kafka is a highly **scalable**, and **distributed** platform for creating and processing streams in real-time.  

> Note: consider of real-time processing of electricity, u will understand what is "real-time" and "streams".  

## How does it work

Kafka adopted pub sub messaging system architecture and it works as an enterprise messaging system.  

- A typical messaging system has got three components:
  - producer: A producer is a client application that sends data records, these data records are called messages.  
  - broker: The broker is responsible for receiving messages from the producers and storing them into local storage.  
  - consumer: the consumers are again client applications that read messages from the broker and process them.  

Kafka works as a pub sub messaging system where we create producer applications to send data as a stream. We install and configure Kafka server to act a message broker. And finally we create consumer applications to process the data stream in real time. That is all.  

Kafka initially started with two things: server software that u can install and configure to work as a message broker. A (Go/Java, etc.) client API library to help with the following: create Kafka producer applications and create consumer applications.  

But later Kafka aspired to become a full fledged real time streaming platform. And to achieve that objective, they augmented Kafka with three more components:  

- Kafka Connect (open source)
- Kafka Streams (open source)
- KSQL (commercial tool)  

So now, from 2011 to 2019, Kafka evolved as a set of five components:  

- Kafka Broker: the central server system.
- Kafka Client: a producer and consumer API library.
- Kafka Connect: which addresses the initial data integration problem for which Kafka was initially designed.
- Kafka Streams: another library for creating real-time stream processing applications.
- KSQL: Kafka is now aiming to becoming a real time database and capture some market share in databases and DWBI space.  

With careful design, the messages can reach from producers to consumers in milliseconds. The producers and consumers are completely decoupled and they do not need tight coupling or direct connections.  

## 