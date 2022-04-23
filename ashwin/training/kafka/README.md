Getting Started with Kafka in Golang
====================================

![](https://miro.medium.com/max/1374/1*6SQiJ4tinE0p4sjoBexxuA.png)

Go and Apache Kafka official logo

This is my experience in these past two weeks while I get my hand dirty in Apache Kafka. Even I was introduced with Kafka by my CTO several months ago, but I still have some problems about how to produce and consume a data to Kafka and just have a spare time to take a look at it again. The harder part that I've encountered was when I try to setup Kafka cluster running using docker. While there are some open-source docker images out there, but they just give an example to running Kafka in single node.

In this post I will tell you briefly about what is Kafka, how to setup it using Docker, and create simple program which produce and consume message from Kafka.

The behind story of this post is about "how to make an API faster to save data, hence make the lower latency and more clients can be served." Also, this post is intended to be used as an internal engineering show at [Qiscus](https://www.qiscus.com/) (yes, we at Qiscus periodically sets the internal sharing to what we've learned so far). So, let's start with what is Apache Kafka itself.

What Is Apache Kafka?
=====================

![](https://miro.medium.com/max/1400/1*iUWxneAQ_kozzLPkFsrakw.png)

Kafka Architecture. Source: <http://kth.diva-portal.org/smash/get/diva2:813137/FULLTEXT01.pdf>

Apache Kafka is an open-source stream processing software platform which started out at Linkedin. It's written in Scala and Java. Kafka is based on commit log, which means Kafka stores a log of records and it will keep a track of what's happening. This [commit log is similar with common RDBMS uses](https://stackoverflow.com/a/45141245/5489910).

To make it easier to imagine, it's more like a queue where you always append a new data into the tail.

![](https://miro.medium.com/max/1238/1*fsjNrvgiqvxmjhuYK6-wLA.png)

How Kafka stores a log, image source: <https://www.confluent.io/blog/okay-store-data-apache-kafka/>

While it seems simple, but in Kafka it can maintain the records into several partitions with same *topic*. Based on Kafka documentation's, a topic is a category or feed name to which records are published. So, rather than just write into one queue like the image above, Kafka can writes into several queue with same topic name.

![](https://miro.medium.com/max/832/1*GoRlq7O8qMNui6tvnq30cg.png)

For each topic, the Kafka cluster maintains a partitioned log that looks like this. Source: <https://kafka.apache.org/documentation/#design>

Imagine a topic is a name of a hotel. And imagine that the partition topic is like that hotel have 3 lift elevators with same direction, let's say: up direction. In a busy day, when a guest of a hotel want to go to upstairs, they will choose one of three elevators which have a few passengers on it. With same analogical logic, when a data arrives, we can tell Kafka to write into specific partition.

It's all about how Kafka handle an incoming data. Now, let's see how we can read or consume the data from Kafka.

In Kafka you can consume data from specific partition of a topic, or you can consume it from all partition. Interestingly, you can subscribe the data using several clients/workers and make each of it retrieve different data from different partition using consumer group. So, what is a consumer group?

Consumer group is like a label of a group of consumer. For each consumer under the same label, you will consume different messages. Hmm okay, in simpler way, it's like when you are at school, your teacher wants you to make a group of 3 persons. Then, your teacher will give each group a label of name, for example group 1 given a name "Tiger", while the other one "Apple", and so on. From now on, you and your other 2 friends are recognized as one entity rather than 3 entity.

![](https://miro.medium.com/max/948/1*J-0xbraSo0fbyrrPCXedlg.png)

Kafka consumer group. Source: <https://kafka.apache.org/documentation/#design>

In Kafka, as the consumer already grouped into one label, it can consume different message, so the job of each person in the same group will be not redundant.