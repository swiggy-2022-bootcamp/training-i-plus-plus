# Online Shopping Application

## Functional Requirements:

### admin 
- admin can register, login and logout
- admin can add products , view products
- admin can delete and update products 
- admin can view all users and individual user

### user management
- user can register, login and logout
- user and view all products and single product 
- user can view only his profile
- user can add products to cart
- user can view items in cart
- user can delete and update items in cart


## Concepts/Topic covered:
- gin-gonic/gin framework
- JWT integration
- Introducing middlewares
- MondoDb integration

## APIs:

### USER 

POST   "users/signup"
POST   "users/login"
GET    "/users"
GET    "/users/:user_id"

### PRODUCT

GET    "/products"
GET    "products/:product_id"
POST   "/products/add"
DELETE "/products/:product_id"
PUT    "/products/:product_id"

### CART

GET    "/cart/all/:user_id"
DELETE "/cart/:cart_id"
POST   "/cart/add"



