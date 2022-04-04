[![Quality Gate Status](http://localhost:8094/api/project_badges/measure?project=Microservies&metric=alert_status&token=d4405061eb691a72d706390a4cece5445665150c)](http://localhost:8094/dashboard?id=Microservies)

This project is a slight modification of Shopping Cart
-------------------------------------------------------
Generally in Shopping Cart User buys different products like Mobile,Bags, etc

But, In this project User buys services provided by employees like Plumber, Painter, Carpenter, Daily Labour, Field Worker

-------------------------------------------------------
</br>
 <img src="https://github.com/swiggy-2022-bootcamp/training-i-plus-plus/tree/main/uday/mini-project/MicroServices_ShoppingCartProject/ServiceProvider/diagram.PNG" width="500" height="500">
 </br>

Definitions:

User   : Customer who buys a service
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
10) Payment after work is completed  //  Since there is no GUI it might be impossible but will try
11) Filter Service Providers based on their rating
12) CRUD of Service Providers in DB
-------------------------------------------------------


API Requests:
=============

1) GET   /expert/service   => to get all the services available
2) GET   /expert/expert?skill=carpenter  =>Automatically CPU assigns an expert to the user
3) POST  /expert/addrating?expertid=3   => to add rating and review to an user
4) GET   /expert/get?skill=carpenter  => it shows all the experts who are carpenters
5) GET   /expert/done?expertid=3   => it releases the expert so system can assign him to other users
6) POST  /expert/signexpert  => creates a new expert 
7) GET   /expert/getexpert?expertid=4  =>get an expert based on his ID
8) GET   /expert/filter?skill=carpenter&rating=4   => filters experts and returns based on the rating
9) POST  /user/signuser  => creates an user
10)POST  /user/loginuser  => checks the user and returns JWT-token
11)GET   /expert/services  => returns all the available services in the system
12)POST  /user/isuserpresent  => returns boolean whether the user is present or not
13)POST  /user/getuser   =>  returns a specfic user
