
This project is a slight modification of Shopping Cart
-------------------------------------------------------
Generally in Shopping Cart User buys different products like Mobile,Bags, etc

But, In this project User buys services provided by employees like Plumber, Painter, Carpenter, Daily Labour, Field Worker

-------------------------------------------------------
Architecture Diagram
====================
</br>

 <img src="https://github.com/swiggy-2022-bootcamp/training-i-plus-plus/blob/main/uday/mini-project/MicroServices_ShoppingCartProject/aserviceProvider/diagram.PNG" width="800" height="500">
</br>

Definitions:

User   : Customer who buys a service</br>
Expert : Person who sells his service (like Carpentry, plumber, painter)  to the User

-------------------------------------------------------
Technologies/tools Used:
=================
    Golang
    gRPC
    SwaggerUI
    Microservices
    smtp.mail
    MongoDB
    Gin
    JWT
    SonarQube
    Kafka
    Docker (for implementing Kafka,SonarQube)
-------------------------------------------------------
Features included in this Project:
=================================
1) Authentication and Authorisation
2) Book Service
3) Rating and Reviewing Service
4) SignUp
5) SignIn through JWT
6) MongoDB for storing Data
7) System : Automatic Expert Suggestion to the User
8) System : Sending Mail to User on Service confirmation
9) Request cancellation and Request Waiting functionalities
10) Payment after work is completed  
11) Filter Service Providers based on their rating
12) CRUD of Service Providers in DB
-------------------------------------------------------


API Requests:
=============

1) GET    /expert/services                    => to get all the services available</br>
2) POST   /expert/addrating/{expertid}        => User adds rating to the service provider
3) GET    /expert/filter/{skill}/{rating}     => Filters the user based on the rating
4) DELETE /expert/delete/{expertid}           => Deletes an expert based on his ID</br>
5) GET    /expert/get/{skill}/{userid}        => Books a server based on the skill and attached userid to it </br>
6) GET    /expert/getallexperts               => Get all experts in database </br>
7) POST   /expert/getexpert/{skill}           => Get Expert by skill </br>
8) GET    /expert/getexpert/{expertid}        => Get an expert based on his ID</br>
9) POST   /user/signupexpert                  => creates a new expert-with service in DB</br>
10) POST  /user/loginuser                     => checks the user and returns JWT-token</br>
11) GET   /expert/waitingreq/{expertid}       => Get the waiting Requests of an Expert  </br>
12) GET   /expert/acceptreq/{expertid}        => Service provider accepts the waiting request of an user
13) GET   /expert/rejectreq/{expertid}        => Service provider rejects the waiting requests of an user</br>
14) GET   /expert/complete/{cost}/{expertid}  => Completes a waiting request and adds the cost to the service </br>
15) POST  /user/loginuser                     =>  Login with username and password => returns JWT</br>
16) POST  /user/signuser                      =>  Sign Up User



==================================
Starting Go User server:</br>
-----------------------------
              go run main.go</br>
Starting go ServiceProvider server</br>
---------------------------------
              go run main.go</br>
              ( I have configured kafka on my system using docker, It might throw an error when you run this server because it tries to find the kafka server and topic in your system )
