package models

import "time"

type ServerStats struct {
    Id                int64         `gorm:"primary_key"`
    ServerId          int64         `gorm:"column:id_server" json:"id"`
    Server            Server        `gorm:"foreignkey:ServerId"`
    SessionId         int64         `gorm:"column:id_session"`
    Session           ServerSession `gorm:"foreignkey:SessionId"`
    Time              *time.Time    `json:"time"`
    Fps               float32       `json:"fps"`
    FpsMax            float32       `json:"fpsMax"`
    FpsMin            float32       `json:"fpsMin"`
    EngineUpdate      float32       `json:"engineUpdate"`
    EngineUpdateMax   float32       `json:"engineUpdateMax"`
    EngineUpdateMin   float32       `json:"engineUpdateMin"`
    ClientUpdate      float32       `json:"clientUpdate"`
    ClientUpdateMax   float32       `json:"clientUpdateMax"`
    ClientUpdateMin   float32       `json:"clientUpdateMin"`
    ShedulerUpdate    float32       `json:"shedulerUpdate"`
    ShedulerUpdateMax float32       `json:"shedulerUpdateMax"`
    ShedulerUpdateMin float32       `json:"shedulerUpdateMin"`
    BandwidthIn       float32       `json:"bandwidthIn"`
    BandwidthOut      float32       `json:"bandwidthOut"`
    ObjectsServer     int           `json:"objectsServer"`
    ObjectsClient     int           `json:"objectsClient"`
    TrafficOut        float32       `json:"trafficOut"`
    TrafficIn         float32       `json:"trafficIn"`
    Players           int           `json:"players"`
}
