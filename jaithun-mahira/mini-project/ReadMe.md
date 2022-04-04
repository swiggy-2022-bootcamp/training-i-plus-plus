# Train Reservation System

Train reservation application where users can book tickets and admins can add train related details
### Features

- User Registration
- User Login
- Authentication using JWT
- Roles Management
- Scheduling Trains
- Reserving/Booking Tickets
- Canceling Tickets
- Logging
- Kafka Producer/Consumer
- Swagger Documentation

### Microservices

- User Microservice

  - User Registration
  - Update User Details
  - User/Admin Login
  - Authentication using JWT
  - Roles Management

- Train Microservice(Admin Operations)

  - Schedule Trains
  - Update/Delete Train Details
  - Search available Trains

- Ticket Microservice(User Operations)

  - Book Tickets
  - Cancel Tickets

### Kafka Implementation Diagram

<img width="1003" alt="Screenshot 2022-04-05 at 12 27 18 AM" src="https://user-images.githubusercontent.com/73777273/161613788-49eca158-d99c-485f-97a4-0cec29fa60d2.png">


### Swagger Documentation
<img width="1422" alt="Screenshot 2022-04-05 at 12 11 15 AM" src="https://user-images.githubusercontent.com/73777273/161613853-39e40c89-014a-4e56-979f-3505778cbaa5.png">

### Postman Collection

Train Service - https://www.postman.com/collections/f49f6a94ec491667bb62

Ticket Service - https://www.getpostman.com/collections/f7d10be53288ed2da54a

User Service - https://www.getpostman.com/collections/f7c89ef3b38b2c1ae659
