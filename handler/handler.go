package handler

import (
	"fmt"
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, saya sedang belajar Golang"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Home dari belajar golang"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, stay calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":"I'm Learning Golang Web",
	// 	"content":"I'm Learning Golang Web",
	// 	"number":90,
	// }

	data := entity.Product{ID: 1, Name: "Mobilio", Price: 22000000, Stock: 3}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, stay calm", http.StatusInternalServerError)
		return
	}
}

func HariHandler(w http.ResponseWriter, r *http.Request) {
	s := "hari page fullstack"
	fmt.Fprintf(w, "%s", s)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}


	// fmt.Fprintf(w, "Product page: %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb, 
	}

	tmpl, err := template.ParseFiles(path.Join("views","product.html"), path.Join("views","layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err !=  nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}