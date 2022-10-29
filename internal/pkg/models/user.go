package models

type User struct {
	BasicModel
	Account  string `gorm:"type:varchar(300);uniqueIndex" json:"account"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Password string `gorm:"type:varchar(600)" json:"password"`
}
