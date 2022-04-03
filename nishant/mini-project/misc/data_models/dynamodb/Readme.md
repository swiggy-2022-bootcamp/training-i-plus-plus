table definitions for dynamoDb


create table : 

```
aws dynamodb --endpoint-url http://localhost:8000  create-table --cli-input-json file://users.json

```