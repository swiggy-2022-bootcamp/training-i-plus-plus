start

```bash
docker-compose up
```


to run console command, we need to bash into kafka container
```sh
docker exec -it [container_name] /bin/sh


docker exec -it w8_kafka_1 /bin/sh

docker exec -it kafka_notification_service_kafka_1 /bin/sh

cd /bin
ls -a
```

```sh
#create topic
./kafka-topics --bootstrap-server localhost:9092 --create --replication-factor 1 --partitions 1 --topic test_topic

#list topic
./kafka-topics --bootstrap-server localhost:9092 --list

#start console consumer
./kafka-console-consumer --bootstrap-server localhost:9092 --topic test_topic

#start console producer
./kafka-console-producer --broker-list localhost:9092 --topic test_topic
```