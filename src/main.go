package main

import (
    "fmt"
    "net/http"
    "coremind/db"
    "coremind/handlers"
)

var ProgressChan = make(chan int)

func init() {
    handlers.ProgressChan = ProgressChan // link ti chanel
}

func StreamProgressHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")

    clientChan := make(chan int)
    defer func() {
        fmt.Println("Клиент отключился")
    }()

    go func() {
        for percent := range clientChan {
            fmt.Fprintf(w, "data: %d\n\n", percent)
            w.(http.Flusher).Flush()
        }
    }()

    for p := range ProgressChan {
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
            StreamProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })
    http.HandleFunc("/upload/dataset", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            handlers.UploadDatasetHandler(w, r)
        } else if r.Method == "GET" {
            StreamProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })
    http.HandleFunc("/upload/code", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            handlers.UploadPyHandler(w, r)
        } else if r.Method == "GET" {
            StreamProgressHandler(w, r)
        } else {
            http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        }
    })


    http.HandleFunc("/logout", handlers.LogoutHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}