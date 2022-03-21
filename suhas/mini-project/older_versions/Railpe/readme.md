# User Mangement with Dynamodb

## Rest api for User Managment

- Create 

    > POST http://localhost:8081/user/add <br>
    > Content-Type: application/json <br>
    >
    >{  <br>
        "Name": "ravi", <br>
        "Email": "ravi@jio.com" <br>
     } <br>

- Update 
    > POST http://localhost:8081/user/update <br>
    > Content-Type: application/json <br>
    >
    >{  <br>
        "UUID": "13579", <br>
        "Email": "ravi@jio.com" <br>
     } <br>

- Read
    > POST http://localhost:8081/user/read <br>
    > Content-Type: application/json <br>
    >
    >{  <br>
        "UUID": "13579", <br>
     } <br>

- Delete
    > POST http://localhost:8081/user/delete <br>
    > Content-Type: application/json <br>
    >
    >{  <br>
        "UUID": "13579", <br>
     } <br>