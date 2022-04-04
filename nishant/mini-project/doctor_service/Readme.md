# Doctor & Appointment Service

- Create Update Delete Doctors
- Get all Doctors
- Get all Appointments of a specific user (by userid)
- Make Appointment for a User. This makes a call to user service for fetching user details
- On creating a Appointment. A msg is published on kafka to send to notification service
- logs are generated in file  /doctor_service/doctor_service.log


## Port = 7451

## Stack
- Gin
- MongoDb
- Swagger http://localhost:7451/swagger/index.html#/



### Generate Swagger Docs
```
/home/neo/go/bin/swag init
```