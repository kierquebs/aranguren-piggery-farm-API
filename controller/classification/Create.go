package classification

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

func Create(c *fiber.Ctx) error {

	newClass := model.ClassificationModel{}
	if err := c.BodyParser(&newClass); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	_, err := database.CCDB.Exec(`INSERT INTO public.t_stock_classification(title,description)
		VALUES ($1,$2);`,
		newClass.Title,
		newClass.Description,
	)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}
	return c.JSON("Success")

}
