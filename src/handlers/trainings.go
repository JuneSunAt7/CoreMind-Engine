package handlers

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

type TrainingParams struct {
    ModelName         string  `json:"model"`
    DatasetName       string  `json:"dataset"`
    Epochs            int     `json:"epochs"`
    BatchSize         int     `json:"batchSize"`
    Optimizer         string  `json:"optimizer"`
    LearningRate      float64 `json:"learningRate"`
    UseAugmentation   bool    `json:"useAugmentation"`
    ShuffleData       bool    `json:"shuffleData"`
}

var TrainingProgressChan chan int

func StartTrainingHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    cookie, err := r.Cookie("session")
    if err != nil || cookie.Value == "" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    username := cookie.Value

    body, _ := io.ReadAll(r.Body)
    var params TrainingParams
    err = json.Unmarshal(body, &params)
    if err != nil {
        http.Error(w, "Ошибка разбора параметров", http.StatusBadRequest)
        return
    }

    modelPath := filepath.Join("uploads", "queue", username, "models", params.ModelName)
    datasetPath := filepath.Join("uploads", "queue", username, "datasets", params.DatasetName)

    if _, err := os.Stat(modelPath); os.IsNotExist(err) {
        http.Error(w, "Файл модели не найден", http.StatusNotFound)
        return
    }
    if _, err := os.Stat(datasetPath); os.IsNotExist(err) {
        http.Error(w, "Файл датасета не найден", http.StatusNotFound)
        return
    }

    go func() {
        for i := 0; i <= 100; i += 5 {
            TrainingProgressChan <- i
            time.Sleep(300 * time.Millisecond)
        }

        cmd := exec.Command("python", "train.py",
            "--model", modelPath,
            "--dataset", datasetPath,
            "--epochs", fmt.Sprintf("%d", params.Epochs),
            "--batch-size", fmt.Sprintf("%d", params.BatchSize),
            "--optimizer", params.Optimizer,
            "--lr", fmt.Sprintf("%f", params.LearningRate),
        )

        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
            fmt.Println("Ошибка выполнения обучения:", err)
        }
    }()

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Обучение начато!")
}