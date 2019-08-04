package models

import (
    "encoding/json"
    "time"
)

type ServerSession struct {
    Id           int64         `gorm:"primary_key" json:"id"`
    DateStart    *time.Time    `gorm:"column:date_start" json:"dateStart"`
    DateStop     *time.Time    `gorm:"column:date_stop" json:"dateStop"`
    ServerId     int64         `gorm:"column:id_server" json:"-"`
    Server       *Server       `gorm:"foreignkey:ServerId" json:"-"`
    Error        string        `json:"error"`
    Status       string        `json:"status"`
    ServerLog    *ServerLog    `gorm:"foreignkey:SessionId" json:"log"`
    ServerReport *ServerReport `gorm:"foreignkey:SessionId" json:"report"`
}

func (d *ServerSession) MarshalJSON() ([]byte, error) {
    type Alias ServerSession

    if d.DateStop == nil {
        return json.Marshal(&struct {
            *Alias
            DateStart int64  `json:"dateStart"`
            DateStop  *int64 `json:"dateStop"`
        }{
            Alias:     (*Alias)(d),
            DateStart: d.DateStart.Unix(),
            DateStop:  nil,
        })
    } else {
        return json.Marshal(&struct {
            *Alias
            DateStart int64 `json:"dateStart"`
            DateStop  int64 `json:"dateStop"`
        }{
            Alias:     (*Alias)(d),
            DateStart: d.DateStart.Unix(),
            DateStop:  d.DateStop.Unix(),
        })
    }

}
