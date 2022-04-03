# Commands

zookeper - `bin/zookeeper-server-start.sh config/zookeeper.properties`
kafka - `bin/kafka-server-start.sh config/server.properties`
consumer - `bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning`
producer - `echo "message to publish" | bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test > /dev/null`
