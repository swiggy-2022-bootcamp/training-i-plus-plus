# Train Reservation System

### Features

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


### Microservices
1. UserService (:8000)
1. TrainService (:8001)
2. TicketService (:8002)

### User Microservice

```
**Routes**
GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
POST   /userRegistration         --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).RegisterUser-fm (3 handlers)
POST   /userLogin                --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).LoginUser-fm (3 handlers)
GET    /user/get/:username       --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).GetUser-fm (4 handlers)
GET    /user/getall              --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).GetAll-fm (4 handlers)
PATCH  /user/update              --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).UpdateUser-fm (4 handlers)
[GIN-debug] DELETE /user/delete/:username    --> github.com/swastiksahoo153/train-reservation-system/controllers.(*UserController).DeleteUser-fm (4 handlers)
```
