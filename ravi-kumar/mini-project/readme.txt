APIs are both unit tested and integration tested

4 micro services :
    REPORTING_SERVICE_SERVER_PORT = 5001
    INVENTORY_SERVICE_SERVER_PORT = 5002
    ORDER_SERVICE_SERVER_PORT = 5003
    USER_SERVICE_SERVER_PORT = 5004

JWT Security - Session, Access and Role Management (checkout API Auth Grid .png image for access control)

InterService Communication - Order service talks with User Service and Inventory Service. All these 3 services also talks to Reporting Service

Kafka - Order Status Monitoring and Server Health Monitoring (checks if server is up or not)

Swagger - Open API 3.0.0 used

Sonar Cube - Test coverage averages to around 85% of statments (uncovered statements are mostly interfaces, structs, etc - 
             which can't be covered)

             (Note: SonarCube files are in service/.scannerwork folder of each service)

Database - MongoDB; Passwords are encrypted before persisting 

JSON utils are abundantly used for type conversion and serialization/deserialization

Logger - Centralized logging mechanism for each service. Gin's default logger is also changed to accomodate multiple writers (STDOUT and .log file)

Exception Handling - Created service specific structs which implements golang's "error" interface that shall act as base error. 
                     Common errors like UnAuthorizedError, MalformedId error, etc are derived from the base struct. Hence no hard
                     coded errors are made on spot.