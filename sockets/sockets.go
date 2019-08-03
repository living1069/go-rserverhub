package sockets

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
    "rserverhub/util"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
    return true
}}

const (
    SERVER_CREATE string = "SERVER_CREATE"
    SERVER_DELETE string = "SERVER_DELETE"
    SERVER_UPDATE string = "SERVER_UPDATE"
    HOST_CREATE   string = "HOST_CREATE"
    HOST_DELETE   string = "HOST_DELETE"
    HOST_UPDATE   string = "HOST_UPDATE"
    HOST_STATS    string = "HOST_STATS"
)

type Channel struct {
    Clients map[*websocket.Conn]bool
}

func NewChannel() *Channel {
    return &Channel{Clients: make(map[*websocket.Conn]bool)}
}

func (p *Channel) Handle(c *gin.Context) {
    ws, e := upgrader.Upgrade(c.Writer, c.Request, nil)
    p.Clients[ws] = true
    defer util.Handle(c)
    util.Check(e)
}

func (p *Channel) Send(obj interface{}) {
    for e := range p.Clients {
        e.WriteJSON(obj)
    }
}

var QueueInfo *Channel = NewChannel()
var QueueStats map[int]*Channel = make(map[int]*Channel)
var QueueLogs map[int]*Channel = make(map[int]*Channel)
