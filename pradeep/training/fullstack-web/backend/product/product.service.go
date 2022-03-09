package product

import (
	"fmt"
	"net/http"
	"strings"
)

const productsPath = "products"

func handleProduct(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", productsPath))
	if len(urlPathSegments[1:]) > 1) {
		w.WriteHeader(http.StatusBasRequest)
		return
	}
	productId, err := strconv.Atoi(urlPathSegments[len])
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handleProducts")
}

func SetUpRoutes(apiBasePath string) {
	productsHandler := http.HandlerFunc(handleProducts)
	productHandler := http.HandlerFunc(handleProduct)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), productsHandler)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), productHandler)
}
