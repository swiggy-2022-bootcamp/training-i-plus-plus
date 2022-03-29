# Tejas - Train Reservation System

## Setup

1. Dev Server
   - gin run main.go
2. Prod Server
   - go build -o app
   - ./app

Database Models

User

- id
- name
- email
- password
- isAdmin

Train

- id
- name
- stations
  - code
  - arrival_time
  - departure_time

Schedule

- date
- trains
  - train_id
  - stations
    - code
    - arrival_time
    - departure_time
  - seats
  - per_station_charge

Reservation

- pnr
- user_id
- train_id
- from_station_code
- to_station_code
- date
- transaction_id
- status
- seat_number

Payment

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

Features

- User Registration
- User Login
- Authentication using JWT
- Roles Management
- Train Reservation
- Train Schedule
- Payment
- Logging

Technologies Used

- Go
- Gin
- MongoDB

[Postman Collection](https://www.postman.com/Rishabh-Mishra/workspace/my-workspace/collection/7084055-85099171-21dd-47d5-8b75-0bd829486b61?action=share&creator=7084055)
