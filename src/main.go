package main

import (
    "fmt"
    "net/http"
    "coremind/db"
    "coremind/handlers"
)

var UploadProgressChan = make(chan int)
var TrainingProgressChan = make(chan int)

func init() {
    handlers.UploadProgressChan = UploadProgressChan
    handlers.TrainingProgressChan = TrainingProgressChan
}

func StreamUploadProgressHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")

    clientChan := make(chan int)
    go func() {
        for p := range clientChan {
            fmt.Fprintf(w, "data: %d\n\n", p)
            w.(http.Flusher).Flush()
        }
    }()

    for p := range UploadProgressChan {
        clientChan <- p
    }
}

func StreamTrainingProgressHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")

    clientChan := make(chan int)
    go func() {
        for p := range clientChan {
            fmt.Fprintf(w, "data: %d\n\n", p)
            w.(http.Flusher).Flush()
        }
    }()

    for p := range TrainingProgressChan {
        clientChan <- p
    }
}

func main() {
    db.InitDB()

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    js := http.FileServer(http.Dir("js"))
    http.Handle("/js/", http.StripPrefix("/js/", js))

    http.HandleFunc("/", handlers.LoginHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/dashboard", handlers.DashboardHandler)

    http.HandleFunc("/upload/model", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            handlers.UploadModelHandler(w, r)
        } else if r.Method == "GET" {
            StreamUploadProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/upload/dataset", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            handlers.UploadDatasetHandler(w, r)
        } else if r.Method == "GET" {
            StreamUploadProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/upload/code", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            handlers.UploadPyHandler(w, r)
        } else if r.Method == "GET" {
            StreamUploadProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/start-training", handlers.StartTrainingHandler)
    http.HandleFunc("/stream/upload", StreamUploadProgressHandler)
    http.HandleFunc("/stream/train", StreamTrainingProgressHandler)
    http.HandleFunc("/api/models", handlers.GetModelsHandler)
    http.HandleFunc("/api/datasets", handlers.GetDatasetsHandler)
    http.HandleFunc("/logout", handlers.LogoutHandler)
    http.HandleFunc("/params", handlers.ParamsHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}