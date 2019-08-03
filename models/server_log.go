package models

type ServerLog struct {
    Id        int64  `gorm:"primary_key" json:"id"`
    ServerId  int64  `gorm:"column:id_server" json:"-"`
    SessionId int64  `gorm:"column:id_session" json:"-"`
    Filename  string `json:"filename"`
}
