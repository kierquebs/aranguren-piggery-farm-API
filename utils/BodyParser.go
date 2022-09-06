package utils

import "github.com/gofiber/fiber/v2"

//BodyParser this will parse the incoming json request body to the model
func BodyParser(c *fiber.Ctx, in interface{}) error {

	err := c.BodyParser(in)
	if err != nil {
		c.Status(500).JSON(fiber.Map{"retCode": 500, "message": err.Error()})
	}
	return err
}
