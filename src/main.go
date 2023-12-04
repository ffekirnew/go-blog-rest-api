package main

import (
	"blog-api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/blogs", handlers.GetAllBlogsHandler).Methods("GET")
	router.HandleFunc("/blogs", handlers.CreateBlogHandler).Methods("POST")
	router.HandleFunc("/blogs/{id}", handlers.GetBlogByIdHandler).Methods("GET")
	router.HandleFunc("/blogs/{id}", handlers.UpdateBlogHandler).Methods("PUT")
	router.HandleFunc("/blogs/{id}", handlers.DeleteBlogHandler).Methods("DELETE")

	fmt.Println("Server listening at: localhost:4040.")
	if err := http.ListenAndServe(":4040", router); err != nil {
		log.Fatal("Server stopped running.")
	}
}
