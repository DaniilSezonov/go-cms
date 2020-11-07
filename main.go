package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerlin/go-cms/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/content-type", api.GetContentTypesHandler).Methods("GET")
	router.HandleFunc("/api/content-type/{id}", api.GetContentTypeById).Methods("GET")
	router.HandleFunc("/api/content-type", api.CreateContentTypeHandler).Methods("POST")
	router.HandleFunc("/api/content-type/{id}", api.UpdateContentTypeHandler).Methods("PUT")
	router.HandleFunc("/api/content-type/{id}", api.DeleteContentTypeHandler).Methods("DELETE")

	fmt.Println("Listen on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
