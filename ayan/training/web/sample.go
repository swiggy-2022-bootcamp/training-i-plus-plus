package main

func main() {
	handler := http.NewServMux()
	handler.handleFunc("/hi", hi)
	http.ListedAndServe(":8080", nil)
}
