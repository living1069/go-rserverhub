package host

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "rserverhub/app"
    "rserverhub/models"
    "rserverhub/sockets"
    "strconv"
)

const StatusUnknown = "UNKNOWN"
const StatusUp = "UNKNOWN"
const StatusRemoving = "REMOVING"

func List(c *gin.Context) {
    var hosts []models.Host
    app.DB.Find(&hosts)
    c.JSON(http.StatusOK, hosts)
}

func One(c *gin.Context) {
    var host models.Host
    app.DB.Where("id = ?", c.Param("id")).Find(&host)
    c.JSON(http.StatusOK, host)
}

func Create(c *gin.Context) {
    var host models.Host
    _ = c.BindJSON(&host)
    host.Status = StatusUnknown
    app.DB.Create(&host)
    c.JSON(http.StatusOK, host)
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.HOST_CREATE, Payload: host})
}

func Delete(c *gin.Context) {
    val, _ := strconv.Atoi(c.Param("id"))
    app.DB.Delete(models.Host{}, "id = ?", val)
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.HOST_DELETE, Payload: val})
}

func Update(c *gin.Context) {
    var host models.Host
    _ = c.BindJSON(&host)
    host.Status = StatusUnknown
    app.DB.Save(&host)
    c.JSON(http.StatusOK, host)
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.HOST_UPDATE, Payload: host})
}

func AgentUpdate(c *gin.Context) {
    type agentUpdateObject struct {
        Host            int64   `json:"host"`
        CPU             float32 `json:"cpu"`
        MemoryTotal     float32 `json:"memoryTotal"`
        MemoryAvailable float32 `json:"memoryAvailable"`
        ServerStatuses  map[int64]string
    }

    var payload agentUpdateObject
    _ = c.BindJSON(&payload)
    var host models.Host
    app.DB.Where("id = ?", payload.Host).Find(&host)

    // TODO: update server statuses
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.HOST_STATS, Payload: payload})
}

func Servers(c *gin.Context) {
    var servers []models.Server
    app.DB.Preload("Configuration").
        Preload("Host").
        Preload("Session", "date_stop is null").
        Preload("Session.ServerLog").
        Preload("Session.ServerReport").
        Where("id_host = ?", c.Param("id")).
        Find(&servers)
    c.JSON(http.StatusOK, servers)
}
