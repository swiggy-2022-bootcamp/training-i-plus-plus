# Train Reservation System

### Features
```
- Register new User
- Login User
- Modify User
- Remove User
- Get List of all the users
- Get user by id

- Register new Train
- Get all the train details
- Get available seats for a train
- Edit Train Information
- Remove Train
- Get train by id

- Book Ticket
- Get all the ticket details
- Get ticket by id
- Update ticket details
- Cancel Ticket

- Swagger UI for all the modules
- Unit Testing and Mocks for all the modules
- Sonarqube for code coverage
```

### Microservices
### UserService (:8000)
![user](https://user-images.githubusercontent.com/53436195/161576435-2305e593-0058-490d-ba35-0ca34b176336.png)

### TrainService (:8001)
![train](https://user-images.githubusercontent.com/53436195/161576487-96d01435-0d0a-463f-9ccb-3cb741613e84.png)

### TicketService (:8002)
![ticket](https://user-images.githubusercontent.com/53436195/161576520-b460a358-de9c-4b6c-877b-4968131afff4.png)

### User Microservice

```
GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
POST   /userRegistration         --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).RegisterUser-fm (3 handlers)
POST   /userLogin                --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).LoginUser-fm (3 handlers)
GET    /user/get/:username       --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).GetUser-fm (4 handlers)
GET    /user/getall              --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).GetAll-fm (4 handlers)
PATCH  /user/update              --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).UpdateUser-fm (4 handlers)
DELETE /user/delete/:username    --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).DeleteUser-fm (4 handlers)
```

### Train Microservice
```
GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
POST   /train/register           --> github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers.(*TrainController).CreateTrain-fm (4 handlers)
GET    /train/get/:train_number  --> github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers.(*TrainController).GetTrain-fm (4 handlers)
GET    /train/getall             --> github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers.(*TrainController).GetAll-fm (4 handlers)
PATCH  /train/update             --> github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers.(*TrainController).UpdateTrain-fm (4 handlers)
DELETE /train/delete/:train_number --> github.com/swastiksahoo153/MicroserviceKafka/TrainModule/controllers.(*TrainController).DeleteTrain-fm (4 handlers)
```

### Ticket Microservice
```
GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
POST   /ticket/book              --> github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers.(*TicketController).CreateTicket-fm (3 handlers)
GET    /ticket/get/:pnr_number   --> github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers.(*TicketController).GetTicket-fm (3 handlers)
GET    /ticket/getall            --> github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers.(*TicketController).GetAll-fm (3 handlers)
PATCH  /ticket/update            --> github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers.(*TicketController).UpdateTicket-fm (3 handlers)
DELETE /ticket/delete/:pnr_number --> github.com/swastiksahoo153/MicroserviceKafka/TicketModule/controllers.(*TicketController).DeleteTicket-fm (3 handlers)
```
