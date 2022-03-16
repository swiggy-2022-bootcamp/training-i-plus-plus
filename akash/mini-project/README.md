# Swiggy ipp Problem Statement
## Online Shopping Cart
#swiggyipp #requirements #todo #golang

## Problem Statement:

Design and develop an online shopping cart web application that enables a buyer to purchase products posted by sellers.

---

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
  
## Sub Modules:

- Customer
- Vendor
- Order
- Review
- Payment
- Notification

---

## Things Done:
 - [x] Requirements narrow down
 - [x] User Data Model 
 - [x] Login Data Model
 - [x] User REST API - Gin + JSON data
 - [x] Mongo DB database
 - [x] Custom Logger
 
