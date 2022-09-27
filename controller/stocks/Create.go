package stocks

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

func Create(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	newStock := model.CreateStockModel{}
	if err := c.BodyParser(&newStock); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	_, err := database.CCDB.Exec(`INSERT INTO public.t_stock(
		added_date,
		last_updated_date,
		initial_weight,
		initial_day_old,
		status)
		VALUES (
			Now() AT TIME ZONE 'Asia/Manila', 
			Now() AT TIME ZONE 'Asia/Manila',
			$1,$2,1);`,
		newStock.Initial_Weight,
		newStock.Initial_Day_Old,
	)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "New stock successfully added."})

}
