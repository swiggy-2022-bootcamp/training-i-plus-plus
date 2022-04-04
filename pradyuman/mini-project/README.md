# Mini Project- Online Shopping Store

## Services

- **Auth**
- **User**
- **Inventory**
- **Cart**
- **Order**

## Auth Service

Using this service a user can sign in by entering his user Info and insert a corresponding enter into the UserDB with a specific role. If the User then logins in with that user info and the credentials are valid, A Jwt is attached to the request with a validated token that contains the role information within. This role information can be used for authorisation decisions.

**Endpoints:**

1. POST `/signup` - Add a new User with the request body containing the user info. The passwords are hashed and stored in the Database
2. POST `/login` - With email and password passed in the Request Body, the user can login. The Jwt returned with the request will have a role attached to it either "BUYER" or "Seller"

## User Service
This service is to manage the User related action.

**Endpoints:**

1. GET `/users` - Get all user info
2. GET `/user/:userId` - Get user info with User Id 
