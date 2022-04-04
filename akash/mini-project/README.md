# Swiggy ipp Problem Statement
## Online Shopping Cart
#swiggyipp #requirements #todo #golang

## Problem Statement:

Design and develop an online shopping cart web application that enables a buyer to purchase products posted by sellers.

---

## Architecture

![Architecture](https://user-images.githubusercontent.com/26066500/161584203-9644d3ab-9c36-459a-bb2b-b0ec21e1153a.png)


## Features

#### Customer/Buyer

- Login (with email and password)
- Search Products by keyword and filter (Pagination)
- See Product detail (name, price, description)
- Read/Write Product rating and review
- View/Add/Remove from Cart
- Make Payment 
- View/Cancel orders (cannot modify after placing)

#### Vendor/Seller

- Login
- Add/Remove/Modify Products
- See Orders for Products

#### Order

- Get Order Detail
- Get Order Status (Placed, InTransit, Delivered, Cancelled)

#### Notification

- Notify Customer about Order status
- Notify Vendor (EOD Summary of Orders)

---
  
## Microservices:

- Customer
- Product
- Order
---

## Things Done:
 - [x] Requirements narrow down
 - [x] Customer Microservice REST API - Gin + JSON + MongoDB 
 - [x] Order Microservice REST API - Gin + JSON + MongoDB
 - [x] Delivery Microservice API - Fibre + JSON + MongoDB
 - [x] Custom Logger
 - [x] Kafka Exchange: Product & Order Service
 - [x] Docker Compose for Infra (Kafka, DBs)
 - [x] Hashed Password
 
