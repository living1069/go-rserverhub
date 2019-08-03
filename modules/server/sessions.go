package server

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "rserverhub/app"
    "rserverhub/models"
    "strconv"
)

func GetSessions(c *gin.Context) {
    var sessions []models.ServerSession
    app.DB.Preload("ServerLog").
        Preload("ServerReport").
        Where("id_server = ?", c.Param("id")).
        Order("date_start DESC").
        Find(&sessions)

    c.JSON(http.StatusOK, sessions)
}

func DeleteSessions(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    app.DB.Delete(models.ServerLog{}, "id_session in (select id from server_sessions where date_stop is not null and id_server = ?)", id)
    app.DB.Delete(models.ServerReport{}, "id_session in (select id from server_sessions where date_stop is not null and id_server = ?)", id)
    app.DB.Delete(models.ServerSession{}, "date_stop is not null and id_server = ?", c.Param("id"))
    c.Status(200)
}
