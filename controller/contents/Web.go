package contents

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

var wcm []model.WebContentModel

func ListAll(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	rows, err := database.CCDB.Query(`SELECT 
					id,
					title,
					value, 
					description
					FROM public.c_web_static
				`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := wcm

	for rows.Next() {
		content := model.WebContentModel{}
		if err := rows.Scan(
			&content.ID,
			&content.Title,
			&content.Value,
			&content.Description,
		); err != nil {
			return err // Exit if we get an error
		}

		// Append stock to result
		result = append(result, content)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}
