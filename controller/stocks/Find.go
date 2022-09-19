package stocks

import (
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"

	"github.com/gofiber/fiber/v2"
)

type qr struct {
	Code string `json:"code"`
}

var vsm []model.ViewStockModel

func FindByQR(c *fiber.Ctx) error {

	qr := new(qr)

	utils.BodyParser(c, qr)

	rows, err := database.CCDB.Query(`SELECT 
					id,
					added_date,
					added_by,
					last_updated_date, 
					updated_by, 
					initial_weight, 
					current_weight,
					type,
					type_description,
					current_price, 
					current_price_last_updated_date,
					remarks
					FROM public.view_t_stock WHERE qr_code = $1`, qr.Code)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
	}

	defer rows.Close()

	result := vsm

	if rows.Next() {
		stock := model.ViewStockModel{}
		if err := rows.Scan(
			&stock.ID,
			&stock.Added_Date,
			&stock.Added_By,
			&stock.Last_Update_Date,
			&stock.Updated_By,
			&stock.Initial_Weight,
			&stock.Current_Weight,
			&stock.Type,
			&stock.Type_Description,
			&stock.Current_Price,
			&stock.Current_Price_Last_Updated_Date,
			&stock.Remarks,
		); err != nil {

			return c.JSON(fiber.Map{"responseCode": 400, "message": "Error: " + err.Error(), "data": result})

		}

		// Append stock to result
		result = append(result, stock)

	} else {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "QR Code not found.", "data": nil})
	}

	// Return Stock in JSON format
	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": result})

}

func ListAll(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	rows, err := database.CCDB.Query(`SELECT 
					id,
					added_date,
					added_by,
					last_updated_date, 
					updated_by, 
					initial_weight, 
					current_weight,
					type,
					type_description,
					current_price, 
					current_price_last_updated_date,
					remarks
					FROM public.view_t_stock
					ORDER BY added_date DESC
				`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := vsm

	for rows.Next() {
		stock := model.ViewStockModel{}
		if err := rows.Scan(
			&stock.ID,
			&stock.Added_Date,
			&stock.Added_By,
			&stock.Last_Update_Date,
			&stock.Updated_By,
			&stock.Initial_Weight,
			&stock.Current_Weight,
			&stock.Type,
			&stock.Type_Description,
			&stock.Current_Price,
			&stock.Current_Price_Last_Updated_Date,
			&stock.Remarks,
		); err != nil {
			return err // Exit if we get an error
		}
		// Append stock to result
		result = append(result, stock)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}
