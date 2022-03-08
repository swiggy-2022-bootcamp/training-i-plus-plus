package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/bhatiachahat/mongoapi/db"
	users "github.com/bhatiachahat/mongoapi/models/users"
)
func insert(user users.User){
	inserted,err:=db.Collection.InsertOne(context.Background(),user)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted user",inserted)
}
func Add(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/x-www-form-urlencode")
newuser := insert()
json.NewEncoder(w).Encode(newuser)
}
// func upadte(userId string){
// 	id,_:=primitive.ObjectIdFromHex(userId)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	filter := bson.M{"_id":id}
// 	upadte := bson.M{"$set":bson.M{"watched":true}}
// }