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
    EngineUpdate      float32       `json:"engineUpdateTime"`
    EngineUpdateMax   float32       `json:"engineUpdateTimeMax"`
    EngineUpdateMin   float32       `json:"engineUpdateTimeMin"`
    ClientUpdate      float32       `json:"clientUpdateTime"`
    ClientUpdateMax   float32       `json:"clientUpdateTimeMax"`
    ClientUpdateMin   float32       `json:"clientUpdateTimeMin"`
    ShedulerUpdate    float32       `json:"shedulerUpdateTime"`
    ShedulerUpdateMax float32       `json:"shedulerUpdateTimeMax"`
    ShedulerUpdateMin float32       `json:"shedulerUpdateTimeMin"`
    BandwidthIn       float32       `json:"bandwidthDownload"`
    BandwidthOut      float32       `json:"bandwidthUpload"`
    ObjectsServer     int           `json:"objectsServer"`
    ObjectsClient     int           `json:"objectsClient"`
    TrafficOut        float32       `json:"trafficOut"`
    TrafficIn         float32       `json:"trafficIn"`
    Players           int           `json:"players"`
}
