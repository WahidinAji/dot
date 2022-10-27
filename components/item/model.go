package item

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	EmployeeId uint   `json:"employee_id" gorm:"column:employee_id"`
	Name       string `json:"name" gorm:"column:name;size:255"`
}

type ItemDeps struct {
	DB *gorm.DB
}
