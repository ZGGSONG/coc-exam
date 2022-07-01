package model

type Subject struct {
	ID         uint64 `gorm:"column:ID;type:integer;primary_key"`
	Type       string `gorm:"column:Type;type:varchar(255);not null"`
	Question   string `gorm:"column:Question;type:varchar(255);not null"`
	Options    string `gorm:"column:Options;type:varchar(255);not null"`
	Answer     string `gorm:"column:Answer;type:varchar(255);not null"`
	UpdateTime string `gorm:"column:UpdateTime;type:time;not null"`
	Remark     string `gorm:"column:Remark;type:varchar(255)"`
}
