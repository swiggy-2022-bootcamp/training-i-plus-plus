package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// user represents data about a record user.
type user struct {
    ID     string  `json:"id"`
    Email  string  `json:"email"`
    Password string  `json:"password"`
  
}

// users slice to seed record user data.
var users = []user{
	{ID: "1", Email: "john@gmail.com", Password: "123456"},
	{ID: "2", Email: "john@gmail.com", Password: "123456"},
	{ID: "3", Email: "john@gmail.com", Password: "123456"},
   }



// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}

// postusers adds an user from JSON received in the request body.
func postUsers(c *gin.Context) {
    var newUser user

    // Call BindJSON to bind the received JSON to
    // newUser.
    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    // Add the new user to the slice.
    users = append(users, newUser)
    c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of users, looking for
    // an user whose ID value matches the parameter.
    for _, a := range users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
    router := gin.Default()
    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserByID)
    router.POST("/users", postUsers)

    router.Run("localhost:8080")
}