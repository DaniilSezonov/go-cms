package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nerlin/go-cms/data"
)

func GetContentTypesHandler(res http.ResponseWriter, req *http.Request) {
	contentTypes := data.GetContentTypes()

	res.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(res)
	encoder.Encode(contentTypes)
}

func GetContentTypeById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["typeID"])
	if err != nil {
		http.Error(res, "Invalid ID.", http.StatusBadRequest)
		return
	}

	contentType, exists := data.GetContentTypeByID(id)
	if !exists {
		http.Error(res, "Content type was not found.", http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(res)
	encoder.Encode(contentType)
}

func CreateContentTypeHandler(res http.ResponseWriter, req *http.Request) {
	contentType := data.ContentType{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&contentType)
	if err != nil {
		http.Error(res, "Invalid JSON.", http.StatusBadRequest)
		return
	}

	data.CreateContentType(&contentType)
	res.WriteHeader(201)

	encoder := json.NewEncoder(res)
	encoder.Encode(contentType)
}

func UpdateContentTypeHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["typeID"])
	if err != nil {
		http.Error(res, "Invalid type ID.", http.StatusBadRequest)
		return
	}

	contentType := data.ContentType{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&contentType)
	if err != nil {
		http.Error(res, "Invalid JSON.", http.StatusBadRequest)
		return
	}

	updated := data.UpdateContentType(id, contentType)
	if updated {
		res.WriteHeader(200)
	} else {
		res.WriteHeader(201)
	}

	encoder := json.NewEncoder(res)
	encoder.Encode(contentType)
}

func DeleteContentTypeHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["typeID"])
	if err != nil {
		http.Error(res, "Invalid ID.", http.StatusBadRequest)
	}

	data.DeleteContentType(id)
	res.WriteHeader(200)
	res.Write([]byte("Deleted."))
}
