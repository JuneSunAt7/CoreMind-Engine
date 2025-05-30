package handlers

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

// Эти каналы теперь разделены
var UploadProgressChan chan int

func UploadModelHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB
    file, handler, err := r.FormFile("model")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    cookie, _ := r.Cookie("session")
    username := cookie.Value

    fmt.Println("Пользователь:", username)

    path := filepath.Join("uploads", "queue", username, "models")
    os.MkdirAll(path, os.ModePerm)

    dstPath := filepath.Join(path, handler.Filename)
    dst, err := os.Create(dstPath)
    if err != nil {
        http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    var total int64
    buf := make([]byte, 32*1024)

    for {
        n, er := file.Read(buf)
        if n > 0 {
            written, _ := dst.Write(buf[:n])
            total += int64(written)
            percent := int((total * 100) / handler.Size)
            UploadProgressChan <- percent
        }
        if er == io.EOF {
            break
        } else if er != nil {
            http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
            return
        }
    }

    UploadProgressChan <- 100
    fmt.Fprintf(w, "Модель %s успешно загружена!", handler.Filename)
}

func UploadDatasetHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB
    file, handler, err := r.FormFile("dataset")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    cookie, _ := r.Cookie("session")
    username := cookie.Value

    fmt.Println("Пользователь:", username)

    path := filepath.Join("uploads", "queue", username, "datasets")
    os.MkdirAll(path, os.ModePerm)

    dstPath := filepath.Join(path, handler.Filename)
    dst, err := os.Create(dstPath)
    if err != nil {
        http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    var total int64
    buf := make([]byte, 32*1024)

    for {
        n, er := file.Read(buf)
        if n > 0 {
            written, _ := dst.Write(buf[:n])
            total += int64(written)
            percent := int((total * 100) / handler.Size)
            UploadProgressChan <- percent
        }
        if er == io.EOF {
            break
        } else if er != nil {
            http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
            return
        }
    }

    UploadProgressChan <- 100
    fmt.Fprintf(w, "Датасет %s успешно загружен!", handler.Filename)
}

func UploadPyHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB
    file, handler, err := r.FormFile("code")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    cookie, _ := r.Cookie("session")
    username := cookie.Value

    fmt.Println("Пользователь:", username)

    path := filepath.Join("uploads", "queue", username, "code")
    os.MkdirAll(path, os.ModePerm)

    dstPath := filepath.Join(path, handler.Filename)
    dst, err := os.Create(dstPath)
    if err != nil {
        http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    var total int64
    buf := make([]byte, 32*1024)

    for {
        n, er := file.Read(buf)
        if n > 0 {
            written, _ := dst.Write(buf[:n])
            total += int64(written)
            percent := int((total * 100) / handler.Size)
            UploadProgressChan <- percent
        }
        if er == io.EOF {
            break
        } else if er != nil {
            http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
            return
        }
    }

    UploadProgressChan <- 100
    fmt.Fprintf(w, "Код %s успешно загружен!", handler.Filename)
}