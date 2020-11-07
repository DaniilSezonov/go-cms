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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID.", 400)
		return
	}

	contentType, exists := data.GetContentTypeByID(id)
	if !exists {
		http.Error(res, "Content type was not found.", 404)
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
		http.Error(res, "Invalid JSON.", 400)
		return
	}

	data.CreateContentType(&contentType)
	res.WriteHeader(201)
	res.Write([]byte("Created."))
}

func UpdateContentTypeHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID.", 400)
		return
	}

	contentType := data.ContentType{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&contentType)
	if err != nil {
		http.Error(res, "Invalid JSON.", 400)
		return
	}

	updated := data.UpdateContentType(id, contentType)
	if updated {
		res.WriteHeader(200)
		res.Write([]byte("Updated."))
	} else {
		res.WriteHeader(201)
		res.Write([]byte("Created."))
	}
}

func DeleteContentTypeHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(res, "Invalid ID.", 400)
	}

	data.DeleteContentType(id)
	res.WriteHeader(204)
	res.Write([]byte("Deleted."))
}
