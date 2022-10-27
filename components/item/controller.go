package item

import (
	"github.com/gofiber/fiber/v2"
)

func (d *ItemDeps) Insert(c *fiber.Ctx) error {
	var in ItemIn
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	res, err := d.InsertRepo(c.Context(), in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "created item successfully!",
		"data":    res,
	})
}

func (d *ItemDeps) GetAll(c *fiber.Ctx) error {
	res, err := d.GetAllRepo(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": res,
	})
}
