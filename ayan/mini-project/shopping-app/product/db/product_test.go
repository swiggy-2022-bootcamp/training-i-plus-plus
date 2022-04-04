package db

// import (
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func TestNewUser(t *testing.T) {

// 	testUser := &User{
// 		Email:    "abc@xyz.com",
// 		Password: "aBc&@=+",
// 		Name:     "Ab Cd",
// 		Address:  "1, Pqr St.",
// 		Zipcode:  951478,
// 		MobileNo: "9874563210",
// 		Role:     "buyer",
// 	}

// 	newUser := NewUser("abc@xyz.com", "aBc&@=+", "Ab Cd", "1, Pqr St.", 951478, "9874563210", "buyer")

// 	assert.EqualValues(t, testUser, newUser)
// }

// func TestSetId(t *testing.T) {

// 	testUser := User{
// 		Email:    "abc@xyz.com",
// 		Password: "aBc&@=+",
// 		Name:     "Ab Cd",
// 		Address:  "1, Pqr St.",
// 		Zipcode:  951478,
// 		MobileNo: "9874563210",
// 		Role:     "buyer",
// 	}

// 	id := primitive.NewObjectID()
// 	testUser.SetId(id)

// 	assert.Equal(t, id, testUser.Id)
// }

// func TestSetCreatedAt(t *testing.T) {

// 	testUser := User{
// 		Email:    "abc@xyz.com",
// 		Password: "aBc&@=+",
// 		Name:     "Ab Cd",
// 		Address:  "1, Pqr St.",
// 		Zipcode:  951478,
// 		MobileNo: "9874563210",
// 		Role:     "buyer",
// 	}

// 	currTime := time.Now()
// 	testUser.SetCreatedAt(currTime)

// 	assert.Equal(t, currTime, testUser.CreatedAt)
// }

// func TestSetUpdatedAt(t *testing.T) {

// 	testUser := User{
// 		Email:    "abc@xyz.com",
// 		Password: "aBc&@=+",
// 		Name:     "Ab Cd",
// 		Address:  "1, Pqr St.",
// 		Zipcode:  951478,
// 		MobileNo: "9874563210",
// 		Role:     "buyer",
// 	}

// 	currTime := time.Now()
// 	testUser.SetUpdatedAt(currTime)

// 	assert.Equal(t, currTime, testUser.UpdatedAt)
// }
