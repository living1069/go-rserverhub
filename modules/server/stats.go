package server

import (
    "github.com/gin-gonic/gin"
    "rserverhub/app"
    "rserverhub/models"
    "time"
)

func SaveStats(c *gin.Context) {
    var stats models.ServerStats
    c.BindJSON(&stats)

    var server models.Server
    app.DB.Preload("Session", "date_stop is null").
        Preload("Session.ServerLog").
        Where("id = ?", stats.ServerId).
        Find(&server)

    if server.Session != nil {
        date := time.Now()
        stats.SessionId = server.Session.Id
        stats.Time = &date
        app.DB.Save(&stats)
    }
}