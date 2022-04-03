### Features Added In this API

- [x] Login with bcrypt hashing.
- [x] Access resources using JWT token 
- [x] Allows to register.
- [x] Allows Admin to add train
- [x] Allows Admin to add available ticket
- [x] Allows User to check available train
- [x] Allows User to Book a ticket
- [x] Logging support
- [x] Use Kafka to produce and consume events
- [x] Created 4 Microservices - Admin,User,Auth and Payment
- [x] Use grpc for remote procdure call
- [x] Added tests
- [x] Added Stripe Payment Gateway(testing)
- [ ] Sonarqube 


## Railyatra

### Features Implemented

<li>User registration for 2 roles - Admin, User Registration
<li>REST APIs to perform CRUD operations on all two roles
<li>Mongodb to persist data
<li>Authorisation using JWT and password hashing
<li>Using kafka to communicate between different microservices
<li>Using Grpc to communicate between different microservices
<li>Integration with Stripe Payment Gateway

### Modules in the application

<li>Admin Module    - PORT 6001
<li>User Module     - PORT 6002
<li>Auth Module     - PORT 6003
<li>Payment Module  - PORT 6012 (This is a grpc server)

<br>

## Admin Service
![Optional Text](new.svg)

<br>

## User Service
![Optiona Text](new1.svg)


<br>

## Auth Service
![Optiona Text](new2.svg)

<br>

## Payment Service
![Optiona Text](new3.svg)
