package stocks

import (
	"log"

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

	//Estimated Weight computation
	finalWeightAvg, err := FinalWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initialWeightAvg, err := InitialWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	qr := new(qr)

	utils.BodyParser(c, qr)

	rows, err := database.CCDB.Query(`SELECT 
					id,
					added_date,
					last_updated_date, 
					initial_weight,
					initial_day_old,
					Now() AT TIME ZONE 'Asia/Manila',
					status
					FROM public.view_t_stock WHERE qr_code = $1 ORDER BY added_date DESC`, qr.Code)
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
			&stock.Last_Update_Date,
			&stock.Initial_Weight,
			&stock.Initial_Day_Old,
			&stock.Current_DateTime,
			&stock.Status,
		); err != nil {
			return err // Exit if we get an error
		}

		//Logic to add status description
		switch stock.Status {
		case 1:
			stock.Status_Description = "In Stock"
		case 2:
			stock.Status_Description = "Sold"
		case 3:
			stock.Status_Description = "Deceased"
		}
		//end

		//------------------------------------------------

		//Age computation
		days, err := utils.CountDays(stock.Added_Date, stock.Current_DateTime)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		stock.Age_By_Days = days + int(stock.Initial_Day_Old)

		if days >= 90 {
			stock.Description = "For Sale"
		} else {
			stock.Description = "Not yet ready to sell"
		}
		//end

		//------------------------------------------------

		var averageAddedWeightPerDay = (finalWeightAvg - initialWeightAvg) / 122

		var estimatedCurrentWeight = stock.Initial_Weight + (averageAddedWeightPerDay * float32(days))

		stock.Estimated_Current_Weight = estimatedCurrentWeight
		//end

		// Append stock to result
		result = append(result, stock)

	} else {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "QR Code not found.", "data": nil})
	}

	// Return Stock in JSON format
	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": result})

}

func ListAll(c *fiber.Ctx) error {

	//Estimated Weight computation
	finalWeightAvg, err := FinalWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initialWeightAvg, err := InitialWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	rows, err := database.CCDB.Query(`SELECT 
					id,
					added_date,
					last_updated_date, 
					initial_weight,
					initial_day_old,
					Now() AT TIME ZONE 'Asia/Manila',
					status
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
			&stock.Last_Update_Date,
			&stock.Initial_Weight,
			&stock.Initial_Day_Old,
			&stock.Current_DateTime,
			&stock.Status,
		); err != nil {
			return err // Exit if we get an error
		}

		//Logic to add status description
		switch stock.Status {
		case 1:
			stock.Status_Description = "In Stock"
		case 2:
			stock.Status_Description = "Sold"
		case 3:
			stock.Status_Description = "Deceased"
		}
		//end

		//------------------------------------------------

		//Age computation
		days, err := utils.CountDays(stock.Added_Date, stock.Current_DateTime)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		stock.Age_By_Days = days + int(stock.Initial_Day_Old)

		if days >= 90 {
			stock.Description = "For Sale"
		} else {
			stock.Description = "Not yet ready to sell"
		}
		//end

		//------------------------------------------------

		var averageAddedWeightPerDay = (finalWeightAvg - initialWeightAvg) / 122

		var estimatedCurrentWeight = stock.Initial_Weight + (averageAddedWeightPerDay * float32(days))

		stock.Estimated_Current_Weight = estimatedCurrentWeight
		//end

		// Append stock to result
		result = append(result, stock)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}

func FinalWeightAvg() (float32, error) {
	var ave float32
	err := database.CCDB.QueryRow("select AVG(final_weight) FROM public.t_final_weight").Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return ave, nil
}

func InitialWeightAvg() (float32, error) {
	var ave float32
	err := database.CCDB.QueryRow("select AVG(initial_weight) FROM public.t_stock WHERE status = 3").Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return ave, nil
}
