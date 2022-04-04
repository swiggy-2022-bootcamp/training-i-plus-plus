
## Online Shopping Application

### Features Implemented

<li>Role-Right Management-registration for 2 roles - Admin, User Registration
<li>REST APIs to perform CRUD operations 
<li>Mongodb to persist data
<li>Authorisation using JWT and password hashing
<li>Using kafka to communicate between different microservices
<li>Swagger implementation of all the microservices
<li>Postman testing and unit testing for all the API's

### API's
   #### Product Module (8080)
   <li> GET("/products")
   <li> POST("/products")
   <li> DELETE("/products/:product_id")
    <li> GET("/products/:product_id")
	<li> PUT("/products/:product_id")
	

   #### User Module (8081)
   <li> POST("/users/signup")
   <li>	POST("/users/login")
   <li> POST("/users/addtocart/:user_id")
   <li> GET("/users/getcart/:user_id")
   <li> GET("/users/:id")
   <li> PUT("/users/:id")
   <li> DELETE("/users/:id")
   <li> GET("/users")

   #### Order Module (8082)
   <li> POST("/orders/place_order/:user_id")

   #### Payment Module (8083)
   <li>  POST("/payment/:order_id")

   #### Track-stream Module (8084)
   <li>	GET ("/getAnalytics")

### Modules in the application

<li>Product Module    - PORT 8080
<li>User Module     - PORT 8081
<li>Order Module     - PORT 8082
<li>Payment Module  - PORT 8083 
<li>Track-Stream Module  - PORT 8084 

<br>

## Product Service
![Optional Text](diagram1.png)

<br>

## Order Service
![Optional Text](diagram2.png)


<br>

## Payment Service
![Optional Text](diagram3.png)

<br>

## Track-stream Service
![Optional Text](diagram4.png)
