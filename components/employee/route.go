package employee

import "github.com/gofiber/fiber/v2"

func (d *EmployeeDeps) EmployeeRoute(route *fiber.App) {
	api := route.Group("/api")
	api.Get("/employees", d.Get)
	api.Post("/employees", d.Insert)
	api.Get("/employees/:id", d.GetById)
	api.Patch("/employees/:id", d.UpdateById)
	api.Delete("/employees/:id", d.DeleteById)
}
