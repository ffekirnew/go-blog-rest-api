package handlers

import (
	"blog-api/config"
	"blog-api/dtos"
	"blog-api/entities"
	"blog-api/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getId(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]

	if !ok {
		return 0, fmt.Errorf("Invalid Id.")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("Invalid Id.")
	}

	return id, nil
}

func GetAllBlogsHandler(w http.ResponseWriter, r *http.Request) {
	blogs := []*dtos.BlogDto{}

	config.DbLock.Lock()
	defer config.DbLock.Unlock()

	for _, blog := range config.Db {
		blogs = append(blogs, &dtos.BlogDto{
			Title:  blog.Title,
			Body:   blog.Body,
			BlogId: blog.BlogId,
		})
	}

	response := &responses.GenericResponse{
		Success: true,
		Message: "Blogs retrieved successfully.",
		Error:   "",
		Value:   blogs,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error occured. Please try again.", http.StatusBadRequest)
	}
}

func GetBlogByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid id.", http.StatusBadRequest)
	}

	config.DbLock.Lock()
	defer config.DbLock.Unlock()

	response := &responses.GenericResponse{
		Success: true,
		Message: "Blog retrieved successfully.",
		Error:   "",
		Value:   config.Db[id],
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error occured. Please try again.", http.StatusBadRequest)
	}
}

func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	request := &dtos.CreateBlogDto{}
	if err := decoder.Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	config.DbLock.Lock()
	config.IdLock.Lock()
	defer config.DbLock.Unlock()
	defer config.IdLock.Unlock()

	config.Db[config.Id] = entities.BlogEntity{
		BlogId: config.Id,
		Title:  request.Title,
		Body:   request.Body,
	}
	config.Id++

	response := &responses.GenericResponse{
		Success: true,
		Message: "Blog created successfully.",
		Error:   "",
		Value:   config.Db[config.Id-1],
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error occured. Please try again.", http.StatusBadRequest)
	}
}

func UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid id.", http.StatusBadRequest)
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	request := &dtos.UpdateBlogDto{}
	if err := decoder.Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	config.DbLock.Lock()
	defer config.DbLock.Unlock()

	config.Db[id] = entities.BlogEntity{
		BlogId: id,
		Title:  request.Title,
		Body:   request.Body,
	}

	response := &responses.GenericResponse{
		Success: true,
		Message: "Blog updated successfully.",
		Error:   "",
		Value:   config.Db[config.Id-1],
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error occured. Please try again.", http.StatusBadRequest)
	}
}

func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid id.", http.StatusBadRequest)
	}

	config.DbLock.Lock()
	defer config.DbLock.Unlock()

	delete(config.Db, id)

	response := &responses.GenericResponse{
		Success: true,
		Message: "Blog deleted successfully.",
		Error:   "",
		Value:   config.Db[config.Id-1],
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Error occured. Please try again.", http.StatusBadRequest)
	}
}
