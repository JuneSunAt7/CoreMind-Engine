package handlers

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"

)

func GetModelsHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session")
    if err != nil || cookie.Value == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    username := cookie.Value
    modelsDir := filepath.Join("uploads", "queue", username, "models")

    files, err := os.ReadDir(modelsDir)
    if err != nil {
        http.Error(w, "Не удалось прочитать папку с моделями", http.StatusInternalServerError)
        return
    }

    var models []string
    for _, file := range files {
        if !file.IsDir() {
            models = append(models, file.Name())
        }
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "[")
    for i, model := range models {
        fmt.Fprintf(w, "\"%s\"", model)
        if i < len(models)-1 {
            fmt.Fprintf(w, ",")
        }
    }
    fmt.Fprintf(w, "]")
}