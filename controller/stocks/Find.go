package stocks

import (
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"

	"github.com/gofiber/fiber/v2"
)

type qr struct {
	Code string `json:"code"`
}

func FindByQR(c *fiber.Ctx) error {

	code := new(qr)

	utils.BodyParser(c, code)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Product Successfully Created",
		"data":    code,
	})
}
