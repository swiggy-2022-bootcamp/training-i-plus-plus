## Zookeeper - Kafka Important Commands For WSL2 UBUNTU & WINDOWS 

## On wsl2-ubuntu

open terminal - 1 
- cd ~
- cd kafka_2.13-2.6.0
- bin/zookeeper-server-start.sh config/zookeeper.properties

open terminal - 2
- cd ~
- cd kafka_2.13-2.6.0
- bin/kafka-server-start.sh config/server.properties

open terminal - 3
- cd ~
- cd kafka_2.13-2.6.0
- bin/kafka-topics.sh --create --topic TrainTicketLelo --bootstrap-server localhost:9092
- bin/kafka-console-producer.sh --topic TrainTicketLelo --bootstrap-server localhost:9092

open terminal - 4
- cd ~
- cd kafka_2.13-2.6.0
- bin/kafka-console-consumer.sh --topic swiggy --from-beginning --bootstrap-server localhost:9092


### --> Go code can access the brokers only when the consumers are already running. So, make sure to start the consumers before trying to run the go code.


At the end
rm -rf /tmp/kafka-logs /tmp/zookeeper

## on Windows

- cd to <c:\kafka> folder and execute the below commands

### starting zookeeper
- .\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties

### starting kafka
- .\bin\windows\kafka-server-start.bat .\config\server.properties

### creating a topic 

- .\bin\windows\kafka-topics.bat --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic TrainTicketLelo

### creating a producer
- .\bin\windows\kafka-console-producer.bat --broker-list localhost:9092 --topic TrainTicketLelo

### creating a consumer
- .\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic TrainTicketLelo -from-beginning
