package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// aboutHandler := func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("About Page"))
	// }

	mux.HandleFunc("/", handler.HomeHandler) //root
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/hari", handler.HariHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Profile"))
	// })

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}


