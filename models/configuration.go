package models

type Configuration struct {
	Id         int64  `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	Executable string `json:"executable"`
	Directory  string `json:"directory"`
	Fsgame     string `json:"fsgame"`
	Arguments  string `json:"arguments"`
	Port       int    `json:"port"`
	Map        string `json:"map"`
}
