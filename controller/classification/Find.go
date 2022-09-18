package classification

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

var cm []model.ClassificationModel

func ListAll(c *fiber.Ctx) error {

	rows, err := database.CCDB.Query(`SELECT 
					id,
					title,
					description
					FROM public.t_stock_classification
				`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := cm

	for rows.Next() {
		class := model.ClassificationModel{}
		if err := rows.Scan(
			&class.ID,
			&class.Title,
			&class.Description,
		); err != nil {
			return err // Exit if we get an error
		}
		// Append stock to result
		result = append(result, class)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}
