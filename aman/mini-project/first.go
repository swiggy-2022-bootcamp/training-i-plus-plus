package main

// func main(){

// 	r.POST("/", test)
// }

// func test(c *gin.Context) {
// 	body := c.Request.Body
// 	decoder := json.NewDecoder(body)
// 	var user1 models.User
// 	err := decoder.Decode(&user1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(user1)
// 	value, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	c.JSON(200, gin.H{
// 		"HI":   value,
// 		"body": user1,
// 	})
// }
