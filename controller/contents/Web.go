package contents

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"
)

var wcm []model.WebContentModel

type updateModel struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

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

func UpdateWebContent(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	update := new(updateModel)
	utils.BodyParser(c, update)

	sqlStatement := `UPDATE public.c_web_static SET value = $2 WHERE title = $1;`
	_, err := database.CCDB.Exec(sqlStatement, update.Title, update.Value)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Updated successfully!"})

}
