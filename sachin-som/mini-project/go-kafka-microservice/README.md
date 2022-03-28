### Project
Online Shopping Cart

### Microservice Strucutures (WIP)
![microservices drawio](https://user-images.githubusercontent.com/64790109/160327506-786fbe7a-7cf9-4666-9d35-98b517d49ae9.png)

### Tech Stack
- GoLang
- Gin Gonic
- MongoDB
- Kafka
- confluentinc/confluent-kafka-go
- segmentio/kafka-go

### Microservices
1. UserMicroService (:8001)
2. InventoryService (:8002)
3. ListingService (:8003)
4. OrderService (:8004)

#### UserMicroService
UserMicroService is responsible to do CRUD operations
related to user entiry.

**Routes**
- POST /v1/users/create
- GET /v1/users/get/:userId

#### InventoryMicroService
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

#### Listing MicroService
ListingMicroService constanlty keep a check on the `products`
topic and consumes the products from `products` topic to store them
to **ProductDB**.
This Service lets the user to show and place order of an specific product.
Once the user placed an order, it publishes the products to `ordered_products`
topic.

**Routes**
- GET v1/listing/show/
- POST v1/listing/place_order/:userId/:productId


#### Order MicroService (* Working)
Order Microservice consumes the products from `ordered_products`
kafka topic and store them to **OrderDB** with `order initiated` status.
This microservie is supposed to communiate with **Billing and Payment
MicroService** to complete the product ordering process.
User can check all the intiated/pending/completed orders using this Service.

**Routes**
- GET /v1/orders/get/:userId
