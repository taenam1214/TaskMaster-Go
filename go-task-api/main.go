package main

import (
    "go-task-api/handlers"
    "go-task-api/middleware"
    "go-task-api/models"
    "go-task-api/utils"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    utils.ConnectDatabase()

    utils.DB.AutoMigrate(&models.User{}, &models.Task{})

    r := mux.NewRouter()

    r.HandleFunc("/register", handlers.Register).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")
    r.Handle("/tasks", middleware.Authenticate(http.HandlerFunc(handlers.GetTasks))).Methods("GET")
    r.Handle("/tasks", middleware.Authenticate(http.HandlerFunc(handlers.CreateTask))).Methods("POST")
    r.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(handlers.UpdateTask))).Methods("PUT")
    r.Handle("/tasks/{id}", middleware.Authenticate(http.HandlerFunc(handlers.DeleteTask))).Methods("DELETE")

    r.HandleFunc("/", handlers.Home).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", r))
}
