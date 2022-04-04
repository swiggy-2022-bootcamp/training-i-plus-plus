Tech Stack Used - 
    - GoLang
    - Gin Gonic
    - Mongo DB
    - JWT
    - Swagger UI 
    - Kafka 

Microservices: Each of these microserverices has their own database and authentication systems. 
    - User Service (Port: 6001)
        This service is provides the registeration and login operations on user entity. All the CRUD operations on users are added in the monolithic version of the project - Link. Login operation provides a session token which needs to be provided later for some of the operations. Unit test is provided in the services. 
        Routes - 
            - POST /user/signup
            - POST /user/login

    - Train Service  (Port: 6002)
        This service provides the option of searching for trains to the end user. Searching for train does not require the user to be logged in to the system.
        It also provides all the CRUD operations on the train entity which can only be done if the user is admin. 
        Routes - 
            POST /train/add
            GET /train/get/:trainnumber
            PUT /train/update/:trainnumber
            DELETE /train/delete/:trainnumber
	
	        GET /search_trains 

    - Reservation Service  (Port: 6003)
        Reservation service provides all the ticket booking related operations to the user. Once the user has booked the seats, "booking" topic sends a message to the notification service via kafka so that the required information can be sent to the user. 
        Routes - 
           POST  /reservation/reserve_ticket
           PUT  /reservation/cancel_reservation
           GET  /reservation/allreservations
           
    - Payment Service (Port: 6004)
        This service is responsible for the payment of the the tickets booked. 
        Routes - 
            POST /payment/pay

    - Notification Service (Port: 6005)
        This service act as a consumer for the all the notification sent by other microservices - reservation and payment 

Kafka brokers - 
    - topic: booking
    - topic: payment
Setup Instructions - 
    Use 'go run main.go' in each of the microservices to run the respective service. 