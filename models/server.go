package models

type Server struct {
    Id              int64          `gorm:"primary_key" json:"id"`
    HostId          int64          `gorm:"column:id_host" json:"-"`
    Host            *Host           `gorm:"foreignkey:HostId" json:"host"`
    AutoRestart     bool           `gorm:"column:autorestart" json:"autorestart"`
    ConfigurationId int64          `gorm:"column:id_configuration" json:"-"`
    Configuration   *Configuration `gorm:"foreignkey:ConfigurationId" json:"configuration"`
    Session         *ServerSession `gorm:"foreignkey:ServerId" json:"session"`

}
