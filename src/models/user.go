package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
}
type UserTrainingResult struct {
    ID             uint   `gorm:"primaryKey"`
    UserID         string
    ModelName      string
    DatasetName    string
    Epochs         int
    BatchSize      int
    Optimizer      string
    LearningRate   float64
    FinalLoss      float64
    FinalAccuracy  float64
}