package stocks

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

func Create(c *fiber.Ctx) error {

	newStock := model.CreateStockModel{}
	if err := c.BodyParser(&newStock); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	_, err := database.CCDB.Exec(`INSERT INTO public.t_stock(
		added_date,
		last_updated_date,
		current_price_last_updated_date,
		added_by,
		updated_by,
		initial_weight,
		current_weight,
		type, 
		current_price,
		remarks)
		VALUES (
			Now() AT TIME ZONE 'Asia/Manila', 
			Now() AT TIME ZONE 'Asia/Manila',
			Now() AT TIME ZONE 'Asia/Manila',
			$1,$2,$3,$4,$5,$6,$7);`,
		newStock.Added_By,
		newStock.Added_By,
		newStock.Initial_Weight,
		newStock.Initial_Weight,
		newStock.Type,
		newStock.Current_Price,
		newStock.Remarks,
	)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "New stock successfully added."})

}
