package contents

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

var wContactModel []model.WebContactModel

func ListContact(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	rows, err := database.CCDB.Query(`SELECT 
					id,
					network,
					mobile_num
					FROM public.c_contact
				`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := wContactModel

	for rows.Next() {
		contact := model.WebContactModel{}
		if err := rows.Scan(
			&contact.ID,
			&contact.Network,
			&contact.Mobile_Num,
		); err != nil {
			return err // Exit if we get an error
		}

		// Append stock to result
		result = append(result, contact)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}
