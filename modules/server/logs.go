package server

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "rserverhub/app"
    "rserverhub/models"
    "rserverhub/sockets"
    "rserverhub/util"
)

type saveRequest struct {
    Id   int64    `json:"id"`
    Logs []string `json:"logs"`
}

func SaveLogs(c *gin.Context) {
    var request saveRequest
    util.Check(c.BindJSON(&request))

    tx := app.DB.Begin()
    defer util.HandleTransaction(c, tx)

    var server models.Server
    app.DB.Preload("Session", "date_stop is null").
        Preload("Session.ServerLog").
        Where("id = ?", request.Id).
        Find(&server)

    if server.Session != nil {
        var file *os.File
        var es error

        if server.Session.ServerLog == nil {
            date := server.Session.DateStart.Format("yyyy_MM_dd-HHmmss")

            var log models.ServerLog
            log.ServerId = server.Id
            log.SessionId = server.Session.Id
            log.Filename = fmt.Sprintf("xray_server_%d_%s.log", server.Id, date)

            tx.Save(&log)

            file, es = os.Create("/storage/logs/" + log.Filename)
            util.Check(es)

        } else {
            file, _ = os.Open(server.Session.ServerLog.Filename)
        }

        for _, e := range request.Logs {
            _, _ = file.WriteString(e)
        }

        _ = file.Close()
    }

    tx.Commit()

    if sockets.QueueLogs[int(server.Id)] != nil {
        sockets.QueueLogs[int(server.Id)].Send(request.Logs)
    }
}

func GetLogs(c *gin.Context) {
    c.JSON(http.StatusOK, []string{})
}
