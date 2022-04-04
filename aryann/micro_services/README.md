### Project
Train Reservation System

### Implemented
- GoLang
- Gin Gonic
- MongoDB
- Kafka
- Swagger UI
- SonarQube
- Logging
- Session Key Management/JWT Tokens
- Encrypting Passwords

### Microservices
1. UserService (:3000)
2. TrainService (:5000)
3. TicketService (:7000)
4. PurchaseService (:8000)

#### UserService
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
TicketService is repsonsible to do CRUD operations related to tickets

**Routes**
- POST /ticket
- GET /ticket/:ticketid
- PUT /ticket/:ticketid
- DELETE /ticket/:ticketid

### PurchaseService
PurchaseService is repsonsible to do CRUD operations related to purchasing tickets

**Routes**
- POST /purchase  
- GET /purchase/:purchaseid
- PUT /purchase/:purchaseid
- DELETE /purchase/:purchaseid