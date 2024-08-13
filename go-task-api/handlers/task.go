package handlers

import (
    "encoding/json"
    "go-task-api/middleware"
    "go-task-api/models"
    "go-task-api/utils"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    json.NewDecoder(r.Body).Decode(&task)

    user := middleware.GetAuthenticatedUser(r)
    task.UserID = user.ID

    utils.DB.Create(&task)
    json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    user := middleware.GetAuthenticatedUser(r)

    var tasks []models.Task
    utils.DB.Where("user_id = ?", user.ID).Find(&tasks)
    json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var task models.Task
    utils.DB.First(&task, id)

    if task.ID == 0 {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    user := middleware.GetAuthenticatedUser(r)
    if task.UserID != user.ID {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    json.NewDecoder(r.Body).Decode(&task)
    utils.DB.Save(&task)
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    var task models.Task
    utils.DB.First(&task, id)

    if task.ID == 0 {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    user := middleware.GetAuthenticatedUser(r)
    if task.UserID != user.ID {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    utils.DB.Delete(&task)
    w.WriteHeader(http.StatusNoContent)
}
