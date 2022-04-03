@REM Script to start kafka service
docker-compose up -d

docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic bookedticket

docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic availticket

docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic train

docker-compose down