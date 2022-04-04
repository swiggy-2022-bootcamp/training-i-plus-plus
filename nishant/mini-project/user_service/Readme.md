# User Service 

- User login signup also delete account
- Get all users
- Passwords are hashed using bcrypt before storing in db
- JWT token based Authentication. users can only update/delete their own account
- Get all appointments of logged in user. This makes a call to doctor service for fetching appointments
- On User SignUp. A msg is published on kafka to send to notification service for sending Welcome message
- logs are generated in file /user_service/user_service.log


## Port = 7450

## Stack
- Gin
- DynamoDb
- Swagger http://localhost:7450/swagger/index.html#/


[DynamoDb Schema for user = users.json](mini-project/misc/data_models/dynamodb/users.json)