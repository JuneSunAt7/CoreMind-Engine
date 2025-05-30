// handlers/training_history.go

package handlers

import (
    "coremind/db"      
    "coremind/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetTrainingHistoryHandler(c *gin.Context) {
    cookie, err := c.Cookie("session")
    if err != nil || cookie == "" {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var results []models.TrainingResult

    db.DB.Where("user_id = ?", cookie).Find(&results)

    c.JSON(http.StatusOK, results)
}