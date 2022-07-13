package model

import "time"

type Subject struct {
	ID         uint64    `gorm:"column:ID;type:integer;primary_key" json:"-"`
	Type       string    `gorm:"column:Type;type:varchar(255);not null" json:"type"`
	Question   string    `gorm:"column:Question;type:varchar(255);not null" json:"question"`
	Options    string    `gorm:"column:Options;type:varchar(255);not null" json:"options"`
	Answer     string    `gorm:"column:Answer;type:varchar(255);not null" json:"answer"`
	UpdateTime time.Time `gorm:"column:UpdateTime;type:time;not null" json:"updateTime"`
	Remark     string    `gorm:"column:Remark;type:varchar(255)" json:"remark"`
}

type CountData struct {
	Type  string `json:"type"`
	Count int64  `json:"count"`
}
