# Train Reservation System

## Implemented
- [x] GoLang
- [x] Gin Gonic as HTTP Web Framework
- [x] MongoDB as Persistent Database
- [x] Kafka for communication between microservices
- [x] Swagger UI
- [x] REST API with Json Response 
- [x] SonarQube with Unit Testing
- [x] Logging in a separate file
- [x] Session Key Management / JWT Tokens
- [x] Error Handling
- [x] Password Encryption using Bcrypt Hashing
- [x] User Registration for 2 roles: User and Admin
- [x] MVC Architecture for each microservice  

## Microservices
1. UserService (:3000)
2. TrainService (:5000)
3. TicketService (:7000)
4. PurchaseService (:8000)

### UserService
UserService is responsible to do authorisation and CRUD operations related to user/admin entry

**Routes**
- POST /register
- POST /login
- GET /user/:userId
- DELETE /user/:userId
- GET /admin/:adminid
- DELETE /admin/:adminid

### TrainService
TrainService is repsonsible to do CRUD operations related to train journeys

**Routes**
- POST /train
- GET /train/:trainid
- PUT /train/:trainid
- DELETE /train/:trainid

### TicketService
TicketService is repsonsible to do CRUD operations related to tickets. It also keeps a check on the `purchased` topic to consume purchased tickets. 

**Routes**
- POST /ticket
- GET /ticket/:ticketid
- PUT /ticket/:ticketid
- DELETE /ticket/:ticketid

### PurchaseService
PurchaseService is repsonsible to do CRUD operations related to purchasing tickets. It also publishes/produces all the purchased tickets to `purchased` kafka topic
to list/consume them in the Ticket Service.

**Routes**
- POST /purchase  
- GET /purchase/:purchaseid
- PUT /purchase/:purchaseid
- DELETE /purchase/:purchaseid
