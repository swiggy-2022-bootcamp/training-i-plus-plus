## TrainTicketLelo - Swiggy IPP Microservices Mini Project

- TrainTicketLelo is a backend-only implementation of a Train Reservation System where the user is allowed to view trains, book & cancel the tickets for the respective trains.
- The Project runs on 4 Microservices - Users Service, Trains Service, Reservations Service, Application Health Service.
- Users Service Runs on Port 8001 and provides all the CRUD related Operations on the Users.
- Reservations Service Runs on Port 8002 and provides Booking & Cancellations on Train Tickets.
- Trains Service Runs on Port 8003 and provides all the CRUD Operations on Trains.
- Application Health Service Runs on Port 8004 and checks for the health of all the above services.

## Features Implemented

- Signup
- Login
- Passwords Stored are Hashed
- Access resources using JWT token
- Role Based Access is implemented. Admin has access to all the APIs while the Traveller has limited Access.
- Allows Admin to add, update, delete Trains
- Allows Traveller to check for all the trains, Book a ticket & Cancel the booked Ticket
- Logging support
- Kafka to produce and consume events
- Error Handling using Custom Errors Package
- Services are health checked every 5 seconds

## Core Technologies Used
1. Golang
2. Gin-Gonic
3. Kafka
4. MongoDB
5. Swagger

## Project Setup

- Place the project folder in src folder of the GOPATH
- Have Kafka producer & consumer running with topic TrainTicketLelo. Refer the [KafkaNotes](./kafkaNotes.md) file For Proper Execution Flow.
- In each of the microserivce folder, run `go run main.go`

## Testing

- All the APIs have been tested on Postman and the collection can be found here: [TrainTicketLelo](https://go.postman.co/workspace/My-Workspace~56369b0b-244a-448f-814c-762e325c0447/collection/20338061-18f32a4d-991b-44ed-8e11-0b586e319bbc?action=share&creator=20338061)

## API Documentation
- Swagger UI has been added to the Users, Trains, Reservations Services and can be viewed on the service port at `/swagger/index.html`
