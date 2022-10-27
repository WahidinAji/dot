package employee

import (
	"context"
	"fmt"
)

func (d *EmployeeDeps) InsertRepo(ctx context.Context, in EmployeeIn) (out EmployeeOut, err error) {
	record := Employee{Name: in.Name}
	result := d.DB.WithContext(ctx).Create(&record)
	if result.Error != nil {
		err = fmt.Errorf("insert into employee failed : %v", result.Error.Error())
		return
	}
	out.Id = record.ID
	out.Name = record.Name
	return
}

func (d *EmployeeDeps) GetAllRepo(ctx context.Context) (out []EmployeeOut, err error) {
	result := d.DB.WithContext(ctx).Select("id", "name").Model(Employee{}).Scan(&out)
	if result.Error != nil {
		err = fmt.Errorf("get all record failed : %v", result.Error.Error())
		return
	}
	return
}

func (d *EmployeeDeps) GetByIdRepo(ctx context.Context, id uint) (out EmployeeOut, err error) {
	var employee Employee
	result := d.DB.WithContext(ctx).Select("id", "name").First(&employee, "id = ?", id).Scan(&out)
	if result.Error != nil {
		err = fmt.Errorf("id not found in employee : %v", result.Error.Error())
		return
	}
	return
}

func (d *EmployeeDeps) UpdateByIdRepo(ctx context.Context, id uint, in EmployeeIn) (out EmployeeOut, err error) {
	var employee Employee
	result := d.DB.WithContext(ctx).Select("id", "name").First(&employee, "id = ?", id)
	if result.Error != nil {
		err = fmt.Errorf("id not found in employee : %v", result.Error.Error())
		return
	}

	result = d.DB.Model(&employee).Update("name", in.Name).Scan(&out)
	if result.Error != nil {
		err = fmt.Errorf("update employee failed : %v", result.Error.Error())
		return
	}
	return
}

func (d *EmployeeDeps) DeleteByIdRepo(ctx context.Context, id uint) (out string, err error) {
	var employee Employee
	result := d.DB.WithContext(ctx).Select("id").First(&employee, "id = ?", id)
	if result.Error != nil {
		err = fmt.Errorf("id not found in employee : %v", result.Error.Error())
		return
	}

	result = d.DB.Delete(&employee)
	if result.Error != nil {
		err = fmt.Errorf("delete employee failed : %v", result.Error.Error())
		return
	}
	out = "deleted successfully!"
	return
}
