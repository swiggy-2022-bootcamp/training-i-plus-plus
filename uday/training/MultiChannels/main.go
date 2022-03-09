package main
import (
	"encoding/json"
)

type course struct{
	Name string `json:"name"` 
}

func main(){
	json.Encoder(course)
}