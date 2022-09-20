package stocks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"
)

func UpdateQR(c *fiber.Ctx, qr string) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	genQR := new(model.GenerateQR)
	utils.BodyParser(c, genQR)

	sqlStatement := `UPDATE public.t_stock SET qr_code = $2 WHERE id = $1;`
	_, err := database.CCDB.Exec(sqlStatement, genQR.ID, qr)
	if err != nil {
		return err
	}

	return nil

}
