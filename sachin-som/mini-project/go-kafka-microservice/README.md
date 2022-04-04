## Project
Online Shopping Cart

## Checklist
- [x] Microservice based architecture, each service running on different port
- [x] Golang
- [x] Used Gin-Gonic as HTTP Web Framework
- [x] JWT Based Authentication
- [x] MongoDB as Persistant Database
- [x] Password encryption
- [x] SwaggerUI for API documentation
- [x] Kafka Implementation
- [x] gRPC Communication
- [x] Mock Testing using mockgen and testify
- [x] Logger - Logging in a separate file 

## Microservice Strucutures
![architecture](https://user-images.githubusercontent.com/64790109/161521331-df5b53c3-4313-4360-ac8e-f2c0c482b5dc.png)

## Tech Stack
- GoLang
- Gin Gonic
- MongoDB
- Kafka
- confluentinc/confluent-kafka-go
- segmentio/kafka-go
- gRPC
- SwaggerUI Docs
- MockGen

## Microservices
1. AuthService (:8000)
1. UserService (:8001)
2. InventoryService (:8002)
3. ListingService (:8003)
4. OrderService (:8004)
5. WalletProtoService(:8005)
6. WalletService (:8006)

### UserMicroService
UserMicroService is responsible to do CRUD operations
related to user entiry.
Currently, This service is Mock tested with 85.6%
code coverage.

**Routes**
- POST /v1/users/create
- GET /v1/users/get/:userId
- PATCH /v1/users/update/:userId
- DELETE /v1/users/delete/:userId

### InventoryMicroService
InvetoryMicroService let the users to create their inventory.
It stores the inventory details to `inventoryDB`. Users can also
add/remove products to their inventory.
Alongwith, this service asyncronously publishes all the products
of an inventory to `products` kafka topic to further list them in
**Listing MicroService.**

**Routes**
- POST v1/inventory/register
- POST v1/inventory/:inventoryId/product/add
- GET  v1/inventory/:inventoryId/product/:productId

### Listing MicroService
ListingMicroService constanlty keep a check on the `products`
topic and consumes the products from `products` topic to store them
to **ProductDB**.
This Service lets the user to show and place order of an specific product.
Once the user placed an order, it publishes the products to `ordered_products`
topic.

**Routes**
- GET v1/listing/show/
- POST v1/listing/place_order/:userId/:productId


### Order MicroService
Order Microservice consumes the products from `ordered_products`
kafka topic and store them to **OrderDB** with `order initiated` status.
This microservice communicates to Wallet Microservice and deducts amount from
user wallet as per product's bill, through gRPCs.
User can check all the intiated/pending/paid/completed orders using this Service.

**Routes**
- GET /v1/orders/get/:userId

### Wallet Microservice
Wallet Microservice consists APIs related to wallet management for user.
This microservice also implements the interface generate by `wallet proto` file.

**Routes**
- POST /v1/wallet/create
- GET /v1/wallet/get/:walletId
- PATCH /v1/wallet/:walletId/add

