package models

import "time"

type ServerStats struct {
	Id                int64         `gorm:"primary_key"`
	ServerId          int64         `gorm:"column:id_server"`
	Server            Server        `gorm:"foreignkey:ServerId"`
	SessionId         int64         `gorm:"column:id_session"`
	Session           ServerSession `gorm:"foreignkey:SessionId"`
	Time              *time.Time
	Fps               float32
	FpsMax            float32
	FpsMin            float32
	EngineUpdate      float32
	EngineUpdateMax   float32
	EngineUpdateMin   float32
	ClientUpdate      float32
	ClientUpdateMax   float32
	ClientUpdateMin   float32
	ShedulerUpdate    float32
	ShedulerUpdateMax float32
	ShedulerUpdateMin float32
	BandwidthIn       float32
	BandwidthOut      float32
	ObjectsServer     int
	ObjectsClient     int
	TrafficOut        float32
	TrafficIn         float32
}
