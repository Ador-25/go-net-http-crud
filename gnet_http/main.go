// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	task_service "net_http/business_logic"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	fmt.Println("API running.............")
	r := mux.NewRouter()
	//go install github.com/swaggo/swag/cmd/swag@latest
	//swag init

	r.HandleFunc("/tasks/all", task_service.GetAllTasks).Methods("GET")
	// GetTaskByID godoc
	// @Summary Get a task by ID
	// @Description Get details of a task by ID
	// @ID get-task-by-id
	// @Produce json
	// @Param id path int true "Task ID"
	// @Success 200 {object} Task
	// @Failure 400 {object} ErrorResponse
	// @Router /tasks/{id} [get]
	r.HandleFunc("/tasks/details", task_service.GetTask).Methods("GET")
	r.HandleFunc("/tasks/add", task_service.AddTask).Methods("POST")
	r.HandleFunc("/tasks/edit", task_service.EditTask).Methods("PUT")
	r.HandleFunc("/tasks/delete", task_service.RemoveTask).Methods("DELETE")
	r.HandleFunc("/test", task_service.Test).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"), // The URL pointing to the Swagger JSON file generated by swag init
	))

	log.Fatal(http.ListenAndServe(":8000", r))

}