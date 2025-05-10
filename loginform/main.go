package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)
// test admin:123456
var users = map[string]string{
    "admin": "$2a$10$QVxrn3ub42ENaS8yze3kleMSCQg.3rqJ8c.jXCBzYqVU.PSU0gKha",
}

func hashPassword(password string) string {
    hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashed)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        expectedPassword, exists := users[username]
        if !exists || bcrypt.CompareHashAndPassword([]byte(expectedPassword), []byte(password)) != nil {
            http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
            return
        }

        http.SetCookie(w, &http.Cookie{
            Name:     "session",
            Value:    username,
            HttpOnly: true,
            Path:     "/",
        })

        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return
    }

    http.ServeFile(w, r, "templates/index.html")
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || cookie.Value == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    http.ServeFile(w, r, "templates/dashboard.html")
}

func uploadModelHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB
    file, handler, err := r.FormFile("model")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    dst, _ := os.Create(filepath.Join("uploads", handler.Filename))
    defer dst.Close()
    io.Copy(dst, file)

    fmt.Fprintf(w, "Модель %s успешно загружена!", handler.Filename)
}

func uploadDatasetHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("dataset")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    dst, _ := os.Create(filepath.Join("uploads", handler.Filename))
    defer dst.Close()
    io.Copy(dst, file)

    fmt.Fprintf(w, "Датасет %s успешно загружен!", handler.Filename)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:     "session",
        Value:    "",
        HttpOnly: true,
        Path:     "/",
        MaxAge:   -1,
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
    os.Mkdir("uploads", os.ModePerm)

    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    js := http.FileServer(http.Dir("js"))
    http.Handle("/js/", http.StripPrefix("/js/", js))

    http.HandleFunc("/", loginHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/dashboard", dashboardHandler)
    http.HandleFunc("/upload/model", uploadModelHandler)
    http.HandleFunc("/upload/dataset", uploadDatasetHandler)
    http.HandleFunc("/logout", logoutHandler)

    fmt.Println("Сервер запущен на http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}