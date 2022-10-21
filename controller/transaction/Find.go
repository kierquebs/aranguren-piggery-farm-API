package transaction

import (
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"

	"github.com/gofiber/fiber/v2"
)

var trnModel []model.ViewTransactionModel

func Find(c *fiber.Ctx) error {

	trnRows, err := database.CCDB.Query(`	SELECT 
											ref_id, 
											trn_date, 
											first_name, 
											middle_name, 
											last_name, 
											mobile_number, 
											price_per_kilo
										FROM public.t_transaction
										GROUP BY 
											ref_id, 
											trn_date, 
											first_name, 
											middle_name, 
											last_name, 
											mobile_number, 
											price_per_kilo
									`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
	}

	defer trnRows.Close()

	result := trnModel

	for trnRows.Next() {
		trn := model.ViewTransactionModel{}
		if err := trnRows.Scan(
			&trn.RefID,
			&trn.Trn_Date,
			&trn.FirstName,
			&trn.MiddleName,
			&trn.LastName,
			&trn.MobileNo,
			&trn.PricePerKilo,
		); err != nil {
			return err // Exit if we get an error
		}

		stockRows, err := database.CCDB.Query(`
		SELECT 
			s.id, 
			s.added_date, 
			s.initial_weight, 
			s.final_weight,
			s.qr_code
		FROM public.t_stock s  
		JOIN public.t_transaction t ON t.stock_id = s.id
		WHERE t.ref_id = '` + trn.RefID + `'
		`)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
		}

		for stockRows.Next() {
			if err := stockRows.Scan(
				&trn.Pigs.ID,
				&trn.Pigs.Added_Date,
				&trn.Pigs.Initial_Weight,
				&trn.Pigs.Final_Weight,
				&trn.Pigs.Qr_Code,
			); err != nil {
				return err // Exit if we get an error
			}

		}

		result = append(result, trn)

	}

	// Return Stock in JSON format
	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": result})

}
