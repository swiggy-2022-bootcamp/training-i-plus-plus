# Mini Project- Online Shopping Store

## Services

- **User** (Port:3000)
- **Inventory** (Port:3001)
- **Cart** (Port:3002)
- **Order** (Port:3003)
- **Auth** (Port:3004)





## Features implemented
- GoLang
- Gin Gonic as HTTP Web Framework
- REST APIs with JSON Response 
- Error Handling
- Unit Tests
- MongoDB as Persistent Database
- Password Encryption
- Custom Logging and Logging middleware
- Session Key Management / JWT Tokens
- Kafka for communication between microservices



## Auth Service

Using this service a user can sign in by entering his user Info and insert a corresponding entry into the UserDB with a specific role. If the User then logins in with that user info and the credentials are valid, A Jwt is attached to the request with a validated token that contains the role information within. This role information can be used for authorization decisions.

**Endpoints:**

1. POST `/signup` - Add a new User with the request body containing the user info. The passwords are hashed and stored in the Database
2. POST `/login` - With the email and password passed in the Request Body, the user can log in. The Jwt returned with the request will have a role attached to it either "BUYER" or "Seller"

## User Service
This service is to manage the User related CRUD action. To access these endpoints you need a valid JWT token

**Endpoints:**

1. GET `/users` - Get all user info
2. GET `/user/:userId` - Get user info with User Id 
3. PUT `/user/:userId` - Update User info
4. DELETE `/user/:userId` - Delete user 

## Inventory Service
This service is to manage the Inventory related CRUD action.
Both "BUYER" and "SELLER" can use the endpoints related to getting the product information but only "SELLER" has the ability to update product entries

**Endpoints:**

- Viewing endpoints 
1. GET `/products` - Get all product info
2. GET `/product/:productId` - Get product info with product Id 

- Update endpoints only for seller 
1. POST `/product` - Add new product info
2. PUT `/product/:productId` - Update product info
3. DELETE `/product/:productId` - Delete product

## Cart Service
This service is to manage the cart-related CRUD action. To access these endpoints you need a valid JWT token 

**Endpoints:**

1. POST `/user` - Add a new item to Cart functionality
2. GET `/cart/:userId` - Get all cart items for a user with specified userId
3. PUT `/cart/:cartId` - Update a cart info with cart ID
4. DELETE `/cart/:cartId` - delete a cart entry
5. GET `placeOrder/:userID` - This endpoint places an order from the cart module. The route handler fetches all the cart entries for that userId and places an order for it. This is achieved by sending messages on the "PlacedOrder" Kafka queue to the Order service.

### OrderService
OrderService is responsible to do CRUD operations related to orders.

**Endpoints**
- GET `/orders/user/:userId`
- GET `/orders/seller/:sellerId`

It also has two functions waiting for messages on two queues. "PlacedOrder" queue for messages from cart service and to add new orders. The other queue "UpdateOrderStatus" is used for updating order status at every checkpoint like "Order processed", "Order shipped", and "Out for delivery" and can be seen as a push notifications functionality.

