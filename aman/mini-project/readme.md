# Online Shopping App

## for Swiggy i++ Mini Project

### ***by Aman Gupta***

## Features

> ### Build-wise
- REST APIs
- JSON 
    - for all communnication
- NoSQL Database - 
    - MongoDB (for persistent storage)
- Encryption 
    - Hashing in storing passwor
- Session handling
    - JWT authentication
    - token expiration 
- Logging
    - All the APIs log information in a   ```.log``` file
- Kafka 
- Swagger UI
    - used gin-swagger and swaggo/files libraries to generate UI config in docs directory
- Sonarqube

> ### Use-wise
- Users


    - Users can Sign Up.

    - Same email or phone number cannot be used to create multiple accounts

    - Users can login and even logout

    - Users have access to linked cart to their account

    - Users can add items to the cart

    - Users can purchase items in the cart easily
    
    - Cart become automatically empty after items from the items are purchased so that users can add items for new order now 
    - Order is generated after purcahsing where user can see purchased items

- Sellers

    - Sellers can Sign Up.

    - Same email or phone number cannot be used to create multiple accounts

    - Sellers can login and even logout

    - Sellers have access to linked inventory to their account

    - Sellers can update stocks of items in their inventory




## Requirements to run in local

- GoLang
- MongoDB
- Kafka

## Setup process

- Clone the repository

- Run go mod tidy

- Start mongoDB in the system and update database url in dbConnection.go

- Start kafka and create a topic and update the topic in producer.go

- Run the project with ```go run main.go```

## APIs

### Users
> Only Users who have made an account and are logged in can access these APIs
- Account
    - POST - /users/signup - User Sign Up
    - POST - /users/login - User Login
    - POST - /users/logout - User Logout
    - GET - /users/:user_id - Get User details
    - POST - /request - Submit request as customer
- Product
    - GET - /products/product_id - Get product details
    - GET - /cart - Get list of all items items in cart
    - POST - /cart/add - Add item to cart
    - POST - /cart/product_id - Remove item from cart
    - POST - /cartUpdate/product_id - Update quantity of item inside cart
    - POST - /cart/purchase - Buy all items in cart
    - POST - /cancelOrder/order_id - Cancel placed order

### Sellers
> Only Sellers who have made an account and are logged in can access these APIs
- Account
    - POST - /sellers/signup - User Sign Up
    - POST - /sellers/login - User Login
    - POST - /sellers/logout - User Logout
    - GET - /sellers/:seller_id - Get User details
    - GET - /request - Get request submitted by customer
- Product
    - POST - /products - Create a new product
    - PATCH - products/product_id - Update product stock units
    - POST - /inventory/seller_id - Get Inventory
    - POST - /orderPaid/order_id - Update payment status of order after receiving payment in case of ```Cash On Deliver```
### Internet
> Anybody on the internet can access these APIs without making an account on the portal or without authenticating
- Product
    - GET - /products - Gives all the items available on the portal with lesser details 