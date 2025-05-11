package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
	"fmt"
)

func UploadModelHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("model")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    dst, _ := os.Create(filepath.Join("../uploads", handler.Filename))
    defer dst.Close()
    io.Copy(dst, file)

    fmt.Fprintf(w, "Модель %s успешно загружена!", handler.Filename)
}

func UploadDatasetHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("dataset")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    dst, _ := os.Create(filepath.Join("../uploads", handler.Filename))
    defer dst.Close()
    io.Copy(dst, file)

    fmt.Fprintf(w, "Датасет %s успешно загружен!", handler.Filename)
}