package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nerlin/go-cms/data"
)

func GetContentByTypeHandler(res http.ResponseWriter, req *http.Request) {
	typeID := req.Context().Value("typeID")
	if typeID == nil {
		http.Error(res, "Invalid type ID.", http.StatusBadRequest)
		return
	}

	contents := data.GetContentByTypeID(typeID.(int))
	result := []data.Content{}
	for content := range contents {
		result = append(result, content)
	}

	res.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(res)
	encoder.Encode(result)
}

func GetContentByIDHandler(res http.ResponseWriter, req *http.Request) {
	id := req.Context().Value("id")
	if id == nil {
		http.Error(res, "Invalid content ID.", http.StatusBadRequest)
		return
	}

	content, exists := data.GetContentByID(id.(int))
	if !exists {
		http.Error(res, "Content was not found.", http.StatusNotFound)
		return
	}

	res.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(res)
	encoder.Encode(content)
}

func CreateContentHandler(res http.ResponseWriter, req *http.Request) {
	typeID := req.Context().Value("typeID")
	if typeID == nil {
		http.Error(res, "Invalid type ID.", http.StatusBadRequest)
		return
	}

	var value interface{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&value)
	if err != nil {
		http.Error(res, "Invalid JSON.", http.StatusBadRequest)
		return
	}

	content := data.Content{TypeID: typeID.(int), Value: value}
	err = data.CreateContent(&content)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(res)
	encoder.Encode(content)
}

func UpdateContentHandler(res http.ResponseWriter, req *http.Request) {
	typeID := req.Context().Value("typeID")
	if typeID == nil {
		http.Error(res, "Invalid type ID.", http.StatusBadRequest)
		return
	}

	id := req.Context().Value("id")
	if id == nil {
		http.Error(res, "Invalid content ID.", http.StatusBadRequest)
		return
	}

	var value interface{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&value)
	if err != nil {
		http.Error(res, "Invalid JSON.", http.StatusBadRequest)
		return
	}

	content := data.Content{ID: id.(int), TypeID: typeID.(int), Value: value}
	updated, err := data.UpdateContent(id.(int), content)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Add("Content-Type", "application/json")
	if updated {
		res.WriteHeader(200)
	} else {
		res.WriteHeader(201)
	}

	encoder := json.NewEncoder(res)
	encoder.Encode(content)
}

func DeleteContentHandler(res http.ResponseWriter, req *http.Request) {
	id := req.Context().Value("id")
	if id == nil {
		http.Error(res, "Invalid content ID.", http.StatusBadRequest)
		return
	}

	data.DeleteContent(id.(int))
	res.WriteHeader(200)
	res.Write([]byte("Deleted."))
}

func ContentMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		vars := mux.Vars(req)
		typeID, err := strconv.Atoi(vars["typeID"])
		if err != nil {
			http.Error(res, "Invalid type ID.", http.StatusBadRequest)
			return
		}
		ctx = context.WithValue(ctx, "typeID", typeID)

		_, detail := vars["id"]
		if detail {
			id, err := strconv.Atoi(vars["id"])
			if err != nil {
				http.Error(res, "Invalid content ID.", http.StatusBadRequest)
				return
			}

			ctx = context.WithValue(ctx, "id", id)
		}

		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
