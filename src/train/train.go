package train

import (
    "coremind/models"
    "encoding/json"
    "fmt"
    "time"

    "gorm.io/gorm"
)

var db *gorm.DB

func ProcessTrainingResult(outputJSON string, username string, params models.TrainingParams) {
    var result map[string]interface{}

    // Используем обычное присваивание =
    var err = json.Unmarshal([]byte(outputJSON), &result)
    
    if err != nil {
        fmt.Println("Ошибка парсинга результата:", err)
        return
    }

    trainingResult := models.TrainingResult{
        UserID:        username,
        ModelName:     params.ModelName,
        DatasetName:   params.DatasetName,
        Epochs:        params.Epochs,
        BatchSize:     params.BatchSize,
        Optimizer:     params.Optimizer,
        LearningRate:  params.LearningRate,
        FinalLoss:     result["final_loss"].(float64),
        FinalAccuracy: result["final_accuracy"].(float64),
        Status:        "completed",
        Timestamp:     time.Now().Unix(),
    }

    db.Create(&trainingResult)
}