package main
import ( 
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)
func maindup(){
	fmt.Println("Hello mod in golang")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",sserveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000",r))
}

func greeter(){
	fmt.Println("Hey there mod users")
}

func sserveHome(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Body)
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}