package models

import "time"

type Host struct {
	Id         int64      `gorm:"primary_key" json:"id"`
	Name       string     `json:"name"`
	Address    string     `json:"address"`
	Status     string     `json:"status"`
	DateUpdate *time.Time `gorm:"column:date_update" json:"date_update"`
	Servers    []*Server  `gorm:"FOREIGNKEY:HostId" json:"servers"`
}
