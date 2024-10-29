package main

import "net/http"

func main() {
	server := http.NewServeMux()

	http.ListenAndServe(":8080", server)
}
