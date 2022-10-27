package item

import (
	// "dot/components/employee"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	EmployeeId uint   `json:"employee_id" gorm:"column:employee_id"`
	Name       string `json:"name" gorm:"column:name;size:255"`
}

type ItemIn struct {
	EmployeeId uint   `json:"employee_id" form:"employee_id"`
	Name       string `json:"name" form:"name"`
}

type ItemOut struct {
	Id uint `json:"id"`
	ItemIn
}


type Employee struct {
	gorm.Model
	Name string `json:"name"`
	Items []Item `json:"items"`
}

//relationship to employee one to one
type ItemList struct{
	gorm.Model
	Name string `json:"name"`
	EmployeeId uint `json:"employee_id"`
	Employee Employee
}

type ItemDeps struct {
	DB *gorm.DB
}
