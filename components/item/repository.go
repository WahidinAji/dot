package item

import (
	"context"
	"dot/components/employee"
	"fmt"
)

func (d *ItemDeps) InsertRepo(ctx context.Context, in ItemIn) (out ItemOut, err error) {
	// fmt.Print(in.EmployeeId)
	tx := d.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		err = fmt.Errorf("begin transaction error : %v", tx.Error.Error())
		return
	}

	defer func() {
		if r := recover(); r != nil {
			errRoll := tx.Rollback().Error
			if errRoll != nil {
				err = fmt.Errorf("rollback error: %v", errRoll.Error())
			}
			err = fmt.Errorf("failed to recovered : %v", r)
		}
	}()

	var employee employee.Employee
	findEmployee := tx.Select("id").First(&employee, "id = ?", in.EmployeeId).Scan(&out.EmployeeId)
	if findEmployee.Error != nil {
		err = fmt.Errorf("cannot find employee : %v", findEmployee.Error.Error())
		return
	}

	record := Item{EmployeeId: in.EmployeeId, Name: in.Name}
	createItem := tx.Create(&record)
	if createItem.Error != nil {
		err = fmt.Errorf("created items failed : %v", createItem.Error.Error())
		return
	}

	out.Id = record.ID
	out.Name = record.Name
	if err = tx.Commit().Error; err != nil {
		return
	}

	return
}

func (d *ItemDeps) GetAllRepo(ctx context.Context) (out []ItemList, err error) {

	db := d.DB.WithContext(ctx)

	var list []ItemList
	db.Preload("Employees").Model(Item{}).Scan(&out)
	fmt.Print(list)

	var employee Employee
	r := db.Preload("Items").Find(&employee)
	fmt.Println(r)
	// db.Select("id","employee_id","name").Model(Item{}).Scan(&out)
	return

}
