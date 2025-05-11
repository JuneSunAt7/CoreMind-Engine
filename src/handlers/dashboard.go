package handlers

import (
    "net/http"
    "fmt"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")

    // check cookie
    if err != nil {
        fmt.Println("Ошибка получения куки:", err)
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    if cookie.Value == "" {
        fmt.Println("Кука пустая")
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    fmt.Println("Пользователь вошёл как:", cookie.Value)

    http.ServeFile(w, r, "../src/templates/dashboard.html")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:     "session",
        Value:    "",
        HttpOnly: true,
        Path:     "/",
        MaxAge:   -1,
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}