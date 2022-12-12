package appointment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
)

func Delete(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	id := c.Params("id")

	sqlStatement := `DELETE  FROM public.t_appointment WHERE id=$1;`
	_, err := database.CCDB.Exec(sqlStatement, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Appointment successfully cancelled."})
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Appointment successfully cancelled."})

}
