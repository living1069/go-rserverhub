package server

import (
    "fmt"
    "github.com/go-resty/resty"
    "rserverhub/sockets"
    "rserverhub/util"
    "time"
)

import (
    "bytes"
    "github.com/gin-gonic/gin"
    "rserverhub/app"
    "rserverhub/models"
)

var httpClient = resty.New()

func init() {
    httpClient.SetTimeout(time.Second * 2)
}

func StartServer(c *gin.Context) {
    tx := app.DB.Begin()
    defer util.HandleTransaction(c, tx)

    var server models.Server
    app.DB.Preload("Session", "date_stop is null").
        Preload("Host").
        Preload("Configuration").
        Where("id = ?", c.Param("id")).Find(&server)

    if server.Session != nil {
        c.JSON(400, response{Message: "Сервер уже запущен"})
        return
    }

    var session models.ServerSession
    date := time.Now()
    session.Status = StatusLoading
    session.ServerId = server.Id
    session.DateStart = &date
    tx.Save(&session)

    util.Check(startServer(&server))
    tx.Commit()
    server.Session = &session
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_UPDATE, Payload: server})
}

func StopServer(c *gin.Context) {
    tx := app.DB.Begin()
    defer util.HandleTransaction(c, tx)

    var server models.Server
    app.DB.Preload("Session", "date_stop is null").
        Preload("Host").
        Preload("Configuration").
        Where("id = ?", c.Param("id")).
        Find(&server)

    if server.Session == nil {
        c.JSON(400, response{Message: "Сервер не запущен"})
        return
    }

    date := time.Now()
    server.Session.Status = StatusDown
    server.Session.DateStop = &date
    tx.Save(server.Session)

    server.Session = nil

    util.Check(stopServer(&server))
    tx.Commit()
    sockets.QueueInfo.Send(sockets.Message{Type: sockets.SERVER_UPDATE, Payload: server})
}

func SendCommand(c *gin.Context) {
    var server models.Server
    app.DB.Preload("Session", "date_stop is null").Where("id = ?", c.Param("id")).
        Preload("Host").
        Find(&server)

    if server.Session == nil {
        c.JSON(400, response{Message: "Сервер не запущен"})
        return
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(c.Request.Body)
    str := buf.String()

    defer util.Handle(c)
    util.Check(sendCommand(&server, str))
}

func startServer(server *models.Server) error {
    uri := fmt.Sprintf("http://%s:9090/api/start/%d", server.Host.Address, server.Id)

    _, err := httpClient.R().SetHeader("Content-Type", "application/json").
        SetBody(server.Configuration).
        Post(uri)

    return err
}

func stopServer(server *models.Server) error {
    uri := fmt.Sprintf("http://%s:9090/api/stop/%d", server.Host.Address, server.Id)

    _, err := httpClient.R().SetHeader("Content-Type", "application/json").
        Post(uri)

    return err
}

func sendCommand(server *models.Server, command string) error {
    uri := fmt.Sprintf("http://%s:9090/api/command/%d", server.Host.Address, server.Id)

    _, err := httpClient.R().SetHeader("Content-Type", "application/json").
        SetBody(command).
        Post(uri)

    return err
}
