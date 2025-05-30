package models

type TrainingResult struct {
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
    Status         string
    Timestamp      int64
}

type TrainingParams struct {
    ModelName       string  `json:"model"`
    DatasetName     string  `json:"dataset"`
    Epochs          int     `json:"epochs"`
    BatchSize       int     `json:"batchSize"`
    Optimizer       string  `json:"optimizer"`
    LearningRate    float64 `json:"learning_rate"`
    UseAugmentation bool    `json:"use_augmentation"`
    ShuffleData     bool    `json:"shuffle_data"`
}