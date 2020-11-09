package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nerlin/go-cms/api"
)

func main() {
	router := mux.NewRouter()
	// ContentType
	router.HandleFunc("/api/content-type", api.GetContentTypesHandler).Methods("GET")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}", api.GetContentTypeById).Methods("GET")
	router.HandleFunc("/api/content-type", api.CreateContentTypeHandler).Methods("POST")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}", api.UpdateContentTypeHandler).Methods("PUT")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}", api.DeleteContentTypeHandler).Methods("DELETE")

	// Content
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}/content", api.GetContentByTypeHandler).Methods("GET")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}/content/{id:[0-9]+}", api.GetContentByIDHandler).Methods("GET")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}/content", api.CreateContentHandler).Methods("POST")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}/content/{id:[0-9]+}", api.UpdateContentHandler).Methods("PUT")
	router.HandleFunc("/api/content-type/{typeID:[0-9]+}/content/{id:[0-9]+}", api.DeleteContentHandler).Methods("DELETE")
	router.Use(api.ContentMiddleware)

	fmt.Println("Listen on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
