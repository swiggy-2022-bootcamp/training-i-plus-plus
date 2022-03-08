https://medium.com/platform-engineer/running-aws-dynamodb-local-with-docker-compose-6f75850aba1e


fake creds works for dynamodb local
https://stackoverflow.com/a/35468780/12613203

list tables
```
aws dynamodb list-tables --endpoint-url http://localhost:8000 --profile ipp
```


create table
```
aws dynamodb --endpoint-url http://localhost:8000 --profile ipp create-table --table-name demo-customer-info --attribute-definitions AttributeName=customerId,AttributeType=S --key-schema AttributeName=customerId,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```


put item
```
aws dynamodb --endpoint-url http://localhost:8000 --profile ipp put-item --table-name demo-customer-info --item '{"customerId":{"S":"1"}, "name":{"S":"one"}}'
```



get item
```
aws dynamodb --endpoint-url http://localhost:8000 --profile ipp get-item --table-name demo-customer-info --key '{"customerId":{"S":"1"}}'
```


scan
```
aws dynamodb --endpoint-url http://localhost:8000 --profile ipp scan --table-name demo-customer-info 
```