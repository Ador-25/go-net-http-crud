// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	task_service "net_http/business_logic"
	sse "net_http/server_sent_events"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("API running.............")
	r := mux.NewRouter()
	r.HandleFunc("/events", sse.EventsHandler)
	r.HandleFunc("/tasks/all", task_service.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/details", task_service.GetTask).Methods("GET")
	r.HandleFunc("/tasks/add", task_service.AddTask).Methods("POST")
	r.HandleFunc("/tasks/edit", task_service.EditTask).Methods("PUT")
	r.HandleFunc("/tasks/delete", task_service.RemoveTask).Methods("DELETE")
	r.HandleFunc("/test", task_service.Test).Methods("GET")
	viewsDir := http.Dir("./views")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(viewsDir)))
	log.Fatal(http.ListenAndServe(":8080", r))
}
