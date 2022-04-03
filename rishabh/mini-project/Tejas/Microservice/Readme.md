# Tejas - Train Reservation System

Features

- User Registration
- User Login
- Authentication using JWT
- Roles Management
- Train Reservation
- Train Schedule
- Payment
- Logging
- Unit Testing

Microservice

- User Microservice

  - User Registration
  - User Login
  - Authentication using JWT
  - Roles Management

- Payment Microservice

  - Payment

- Reservation Microservice

  - Train Reservation
  - Train Schedule

- Train Microservice

  - Train Listing

Database Models

- User

  - id
  - name
  - email
  - password
  - isAdmin

- Train

  - id
  - name
  - stations
    - code
  - status

- Schedule

  - date
  - trains

    - train_id
    - stations
      - code
      - arrival_time
      - departure_time
    - seats
    - per_station_charge

- Reservation

  - pnr
  - user_id
  - train_id
  - from_station_code
  - to_station_code
  - date
  - transaction_id
  - status
  - seat_number

- Payment

  - transaction_id
  - user_id
  - amount
  - status

APIs Exposed

- /api/users/register
- /api/users/login

- /api/train/add
- /api/train/remove

- /api/schedule/
- /api/schedule/reserve
- /api/schedule/add

- /api/payment/pay

Technologies Used

- Go
- Gin
- MongoDB

## Setup

1. zookeper - `bin/zookeeper-server-start.sh config/zookeeper.properties`
2. kafka - `bin/kafka-server-start.sh config/server.properties`
3. Run `go run main.go` in each microservice

[Postman Collection](https://www.postman.com/Rishabh-Mishra/workspace/my-workspace/collection/7084055-30eb72c9-42c6-43b3-b20c-d68dbc416e11?action=share&creator=7084055)
