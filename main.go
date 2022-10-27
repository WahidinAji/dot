package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"dot/components/employee"
	"dot/components/item"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn, ok := os.LookupEnv("MYSQL_URL")
	if !ok {
		log.Printf("you need to set MYSQL_URL environment variable")
	}
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("error opening database %v", err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("error connection database with gorm %v", err)
	}

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Printf("error open bigcache: %v", err)
	}

	err = cache.Set("anu", []byte("foo"))
	if err != nil {
		log.Printf("error set key-unique bigcache: %v", err)
	}

	entry, err := cache.Get("anu")
	if err != nil {
		log.Printf("error get key-unique bigcache: %v", err)
	}
	fmt.Println(string(entry))

	err = Migration(gormDB)
	if err != nil {
		log.Printf("error migration database with gorm %v", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	employee := employee.EmployeeDeps{DB: gormDB}
	employee.EmployeeRoute(app)

	app.Listen(":3000")
}

func Migration(db *gorm.DB) error {

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return fmt.Errorf("error")
	}

	ifEmployeeExists := tx.Migrator().HasTable(&employee.Employee{})
	ifItemExists := tx.Migrator().HasTable(&item.Item{})
	if ifEmployeeExists && ifItemExists {
		log.Printf("Table employees and items already migrate!")
		return nil
	}

	//create table
	// err := tx.Migrator().CreateTable(&employee.Employee{})
	// if err != nil {
	// 	log.Printf("migrate employees table failed: %v", err)
	// 	tx.Rollback()
	// 	return err
	// }
	// err = tx.Migrator().CreateTable(&item.Item{})
	// if err != nil {
	// 	errRoll := tx.Rollback().Error
	// 	if errRoll != nil {
	// 		return fmt.Errorf("rollback on failed to migrate items table failed: %v", errRoll)
	// 	}
	// 	return fmt.Errorf("migrate items table failed: %v", err)
	// }

	//migrate table and column
	err := tx.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().AutoMigrate(&employee.Employee{}, &item.Item{})
	if err != nil {
		errRoll := tx.Rollback().Error
		if errRoll != nil {
			return fmt.Errorf("rollback on failed to migrate table: %v", errRoll)
		}
		return fmt.Errorf("migrate table failed: %v", err)
	}

	//create relationship between items.employee_id and employees.id
	err = tx.Migrator().CreateConstraint(&employee.Employee{}, "Items")
	if err != nil {
		errRoll := tx.Rollback().Error
		if errRoll != nil {
			return fmt.Errorf("rollback on failed to create database foreign key for employees & items: %v", errRoll)
		}
		return fmt.Errorf("create database foreign key for employees & items failed: %v", err)
	}
	err = tx.Migrator().CreateConstraint(&employee.Employee{}, "fk_employees_items")
	if err != nil {
		errRoll := tx.Rollback().Error
		if errRoll != nil {
			return fmt.Errorf("rollback on failed to migrate items column: %v", errRoll)
		}
		return fmt.Errorf("migrate items column failed: %v", err)
	}

	err = tx.Commit().Error
	if err != nil {
		errRoll := tx.Rollback().Error
		if errRoll != nil {
			return fmt.Errorf("rollback on failed to commit transaction: %v", errRoll)
		}
		return fmt.Errorf("commit transaction failed: %v", err)
	}
	log.Printf("Migration success!")
	return nil
}
