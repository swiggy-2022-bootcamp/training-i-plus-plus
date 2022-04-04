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
```

![user](https://user-images.githubusercontent.com/53436195/161572838-538e502b-7560-4efa-a260-cd9a5b08cf51.png)

![ticket](https://user-images.githubusercontent.com/53436195/161572884-2c18675f-db03-4389-98a0-0f3dd30a7112.png)

![train](https://user-images.githubusercontent.com/53436195/161572908-da0e45b6-eb25-48e5-a3fb-0aac3b585c90.png)


### Microservices
1. UserService (:8000)
1. TrainService (:8001)
2. TicketService (:8002)

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
