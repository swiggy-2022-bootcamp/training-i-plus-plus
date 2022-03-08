package user
import (
	"github.com/gorilla/mux"
	users "github.com/bhatiachahat/mongoapi/controllers/users"
)
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/add",users.Insert).Methods("POST")
	return router
}
