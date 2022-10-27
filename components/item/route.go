package item

import "github.com/gofiber/fiber/v2"

func (d *ItemDeps) ItemRoute(route *fiber.App) {
	api := route.Group("/api")
	api.Get("/items", d.GetAll)
	api.Post("/items", d.Insert)
}
