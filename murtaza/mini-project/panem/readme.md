# User Management Service

> Panem: named after a democratic constitutional republic capital city in the Hunger Games

### Responsibilities

1. User Authentication
   * Generates JWT token
2. User Authorization
   * Validates JWT token, Decrypts JWT token using secret
   * returns userId and role by decrypting JWT token
3. User Profile management
   * CRUD operations
   * Sign Up / Log In 
   * Update User 
   * Kafka Consumer to update Purchase History
   * /auth API to validate auth-token
