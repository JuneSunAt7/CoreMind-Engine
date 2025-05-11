package handlers

import (
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "coremind/models"
	"coremind/db"
)

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        var user models.User
        if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
            http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
            return
        }

        if !checkPasswordHash(user.Password, password) {
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

    http.ServeFile(w, r, "../src/templates/index.html")
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Проверка на пустые поля
        if username == "" || password == "" {
            http.Error(w, "Все поля обязательны", http.StatusBadRequest)
            return
        }
		hashed, _ := hashPassword(password)

        user := models.User{
            Username: username,
            Password: hashed,
        }

        result := db.DB.Create(&user)
        if result.Error != nil {
            http.Error(w, "Пользователь с таким именем уже существует", http.StatusConflict)
            return
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return
    }

    http.ServeFile(w, r, "../src/templates/register.html")
}