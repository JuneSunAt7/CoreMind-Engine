package main
import (
    "fmt"
    "net/http"
    "coremind/db"
    "coremind/handlers"
    "log"
)
func main() {

    db.InitDB()
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    js := http.FileServer(http.Dir("js"))
    http.Handle("/js/", http.StripPrefix("/js/", js))


    http.HandleFunc("/", handlers.LoginHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/dashboard", handlers.DashboardHandler)
    http.HandleFunc("/upload/model", handlers.UploadModelHandler)
    http.HandleFunc("/upload/dataset", handlers.UploadDatasetHandler)
    http.HandleFunc("/logout", handlers.LogoutHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)

    fmt.Println("Запуск HTTPS сервера на https://localhost:8080")
    err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
    if err != nil {
        log.Fatalf("Ошибка запуска HTTPS сервера: %v", err)
    }
}