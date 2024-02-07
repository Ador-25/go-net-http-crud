package task_service

import (
	"fmt"
	"net/http"
	"net_http/data_access"
	"net_http/models"
	"net_http/utils"
	"strconv"
)

func GetAllTasks(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("function called")
	data, err := data_access.GetAllTasks()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error retrieving tasks"))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("function called")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Invalid ID"))
		return
	}
	data, err := data_access.GetTask(id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error retrieving tasks"))
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func AddTask(rw http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := utils.DecodeJSONBody(r, &task)
	if err != nil {
		http.Error(rw, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	data_access.AddTask(task)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("Task saved successfully"))
}

func EditTask(rw http.ResponseWriter, r *http.Request) {
	var task models.Task
	var id int
	id, err := utils.ParseInt(r)
	err = utils.DecodeJSONBody(r, &task)
	if err != nil {
		http.Error(rw, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	data_access.EditTask(id, task)
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte("Task saved successfully"))
}

func RemoveTask(rw http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseInt(r)
	if err != nil {
		http.Error(rw, "ID NOT FOUND "+err.Error(), http.StatusBadRequest)
		return
	}
	data_access.RemoveTask(id)
}

func Test(rw http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Test function called")
	rw.Write([]byte(message))
}
