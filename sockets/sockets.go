package sockets

import (
    "encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "log"
    "net/http"
    "os"
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
    Name    string
}

func NewChannel(name string) *Channel {
    return &Channel{Clients: make(map[*websocket.Conn]bool), Name: name}
}

func (p *Channel) Handle(c *gin.Context) {
    ws, e := upgrader.Upgrade(c.Writer, c.Request, nil)
    p.Clients[ws] = true
    defer util.Handle(c)
    util.Check(e)
}

func (p *Channel) Send(obj interface{}) {
    if _, b := os.LookupEnv("SOCKET_LOG"); b {
        v, _ := json.Marshal(obj)
        log.Printf("%s = %s", p.Name, string(v))
    }

    for e := range p.Clients {
        e.WriteJSON(obj)
    }
}

var QueueInfo *Channel = NewChannel("/queue/info")
var QueueStats map[int]*Channel = make(map[int]*Channel)
var QueueLogs map[int]*Channel = make(map[int]*Channel)
