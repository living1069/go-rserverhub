package server

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "rserverhub/app"
    "rserverhub/models"
    "rserverhub/sockets"
    "strconv"
    "time"
)

type editRequest struct {
    Id            int64                `json:"id"`
    Host          models.Host          `json:"host"`
    Configuration models.Configuration `json:"configuration"`
    AutoRestart   bool                 `json:"autorestart"`
}

type agentUpdateRequest struct {
    Server int64  `json:"server"`
    Error  string `json:"error"`
    Status string `json:"status"`
}

type response struct {
    Message string `json:"message"`
}

const (
    StatusLoading = "LOADING"
    StatusUp      = "UP"
    StatusDown    = "DOWN"
    StatusError   = "ERROR"
)

func Get(c *gin.Context) {
    var server models.Server
    app.DB.Preload("Configuration").
        Preload("Host").
        Preload("Session", "date_stop is null").
        Preload("Session.ServerLog").
        Preload("Session.ServerReport").
        Where("id = ?", c.Param("id")).
        Find(&server)
    c.JSON(http.StatusOK, server)
}

func Update(c *gin.Context) {
    var request editRequest
    _ = c.BindJSON(&request)

    app.DB.Save(&request.Configuration)

    var server models.Server
    server.Id = request.Id
    server.HostId = request.Host.Id
    server.ConfigurationId = request.Configuration.Id
    server.AutoRestart = request.AutoRestart
    app.DB.Save(server)

    server.Configuration = &request.Configuration

    if request.Id != 0 {
        sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_UPDATE, Payload: server})
    } else {
        sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_CREATE, Payload: server})
    }

}

func Delete(c *gin.Context) {
    val, _ := strconv.Atoi(c.Param("id"))
    app.DB.Delete(models.Server{}, "id = ?", val)
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_DELETE, Payload: val})
}

func All(c *gin.Context) {
    var servers []models.Server
    app.DB.Preload("Configuration").
        Preload("Host").
        Preload("Session", "date_stop is null").
        Preload("Session.ServerLog").
        Preload("Session.ServerReport").
        Find(&servers)

    c.JSON(http.StatusOK, servers)
}

func AgentUpdate(c *gin.Context) {
    var request agentUpdateRequest
    c.BindJSON(&request)

    var server models.Server
    app.DB.Preload("Session", "date_stop is null").
        Where("id = ?", request.Server).
        Find(&server)

    if server.Session != nil {
        if request.Status == StatusError || request.Status == StatusDown {
            server.Session.DateStop = time.Now()
            server.Session.Error = request.Error
        }

        server.Session.Status = request.Status
        app.DB.Save(server.Session)

        sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_UPDATE, Payload: server})
        if request.Status == StatusError && server.AutoRestart {
            startServer(&server)
        }

        c.Status(200)
    } else {
        c.JSON(400, response{Message: "Сервер не запущен"})
    }
}
