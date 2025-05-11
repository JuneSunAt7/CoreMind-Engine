package main
import (
    "fmt"
    "net/http"
    "coremind/db"
    "coremind/handlers"
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

    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}