# Apache Kafka

Apache Kafka is a community distributed event streaming platform capable of handling trillions of events a day. It is a distributed event store and stream-processing platform. It is an open-source system developed by the Apache Software Foundation written in Java and Scala. The project aims to provide a unified, high-throughput, low-latency platform for handling real-time data feeds.

## Model

A pub/sub model allows messages to be broadcasted asynchronously across multiple sections of the applications. The core component that facilitates this functionality is something called a Topic. The publisher will push messages to a Topic, and the Topic will instantly push the message to all the subscribers.

Kafka offers a Pub-sub and queue-based messaging system. Moreover, producers send the message to a topic and the consumer can select any one of the message systems according to their wish.

## Benefits

- Highly scalable. Kafka is a distributed system, which is able to be scaled quickly and easily without incurring any downtime. Apache Kafka is able to handle many terabytes of data without incurring much at all in the way of overhead.
- Highly durable. Kafka persists the messages on the disks, which provides intra-cluster replication. This makes for a highly durable messaging system.
- Highly Reliable. Kafka replicates data and is able to support multiple subscribers. Additionally, it automatically balances consumers in the event of failure. That means that it’s more reliable than similar messaging services available.
- High Performance. Kafka delivers high throughput for both publishing and subscribing, utilizing disk structures that are capable of offering constant levels of performance, even when dealing with many terabytes of stored messages.