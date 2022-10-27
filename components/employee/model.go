package employee

import (
	"github.com/allegro/bigcache/v3"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name;size:255"`
}

type EmployeeDeps struct {
	DB    *gorm.DB
	Cache *bigcache.BigCache
}
type EmployeeIn struct {
	Name string `json:"name" form:"name"`
}
type EmployeeOut struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
