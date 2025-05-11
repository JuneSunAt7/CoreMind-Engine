package handlers

import (
    "net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || cookie.Value == "" {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    http.ServeFile(w, r, "../templates/dashboard.html")
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