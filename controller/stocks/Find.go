package stocks

import (
	"fmt"
	"log"

	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"

	"github.com/gofiber/fiber/v2"
)

type GeneralExpectedWeightModel struct {
	Avg_Added_Date           string  `json:"avg_added_date"`
	Avg_Initial_Weight       float32 `json:"avg_initial_weight"`
	Avg_Final_Weight         float32 `json:"avg_final_weight"`
	Avg_Added_Weight_Per_Day float32 `json:"avg_added_weight_per_day"`
	Days                     int     `json:"days"`
	Estimated_Weight         float32 `json:"estimated_weight"`
	Projected_Weight         []Projected_Weight
}

type Projected_Weight struct {
	Code   int     `json:"code"`
	Weight float32 `json:"weight"`
}

type qr struct {
	Code string `json:"code"`
}

var vsm []model.ViewStockModel

func FindByQR(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
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

		if stock.Age_By_Days >= 120 {
			stock.Description = "Ready to sell"
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

func FindByID(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	id := c.Params("id")
	//Estimated Weight computation
	finalWeightAvg, err := FinalWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initialWeightAvg, err := InitialWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	rows, err := database.CCDB.Query(`SELECT 
					id,
					added_date,
					last_updated_date, 
					initial_weight,
					initial_day_old,
					Now() AT TIME ZONE 'Asia/Manila',
					status
					FROM public.view_t_stock WHERE id = $1 `, id)
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

		stock.Estimated_Current_Weight = estimatedCurrentWeight + stock.Initial_Weight
		//end

		// Append stock to result
		result = append(result, stock)

	} else {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "ID not found.", "data": nil})
	}

	// Return Stock in JSON format
	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": result})

}

func ListAll(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	//Estimated Weight computation
	finalWeightAvg, err := FinalWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initialWeightAvg, err := InitialWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println(fmt.Sprintf("--------------- \n  Average Final Weight: %v \n Average Initial Weight: %v", finalWeightAvg, initialWeightAvg))

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
					WHERE status != 2 AND status != 3
					ORDER BY id DESC
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

		if stock.Age_By_Days >= 120 {
			stock.Description = "Ready to sell"
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
	err := database.CCDB.QueryRow("select AVG(final_weight) FROM public.t_stock WHERE status = 2").Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return ave, nil
}

func InitialWeightAvg() (float32, error) {
	var ave float32
	err := database.CCDB.QueryRow(`	SELECT AVG(t.initial_weight) 
									FROM public.t_stock t
									WHERE status = 2
									`).Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return ave, nil
}

func InitialDayOldAvg() (float32, error) {
	var ave float32
	err := database.CCDB.QueryRow(`	SELECT AVG(t.initial_day_old) 
									FROM public.t_stock t
									WHERE status = 2
									`).Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return ave, nil
}

func GetCurrentDate() (string, error) {
	var now string
	err := database.CCDB.QueryRow(`SELECT Now() AT TIME ZONE 'Asia/Manila'`).Scan(&now)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return now, nil
}

func AddedDateAvg() (string, error) {
	var ave string
	err := database.CCDB.QueryRow(`with
			__ts as(
				select 
				added_date as ts
				FROM public.t_stock
				WHERE status = 2
			)
		select
			to_timestamp(sum(extract(epoch from ts)) / (select count(1) from __ts))
		from
			__ts`).Scan(&ave)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return ave, nil
}

func GeneralExpectedWeight(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	compute := GeneralExpectedWeightModel{}
	//Age computation

	aveDateAdded, err := AddedDateAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	dateNow, err := GetCurrentDate()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	days, err := utils.CountDays(aveDateAdded, dateNow)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	//Estimated Weight computation
	finalWeightAvg, err := FinalWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	initialWeightAvg, err := InitialWeightAvg()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var averageAddedWeightPerDay = (finalWeightAvg - initialWeightAvg) / 122

	var estimatedCurrentWeight = initialWeightAvg + (averageAddedWeightPerDay * float32(days))

	compute.Avg_Added_Date = aveDateAdded
	compute.Avg_Added_Weight_Per_Day = averageAddedWeightPerDay
	compute.Avg_Final_Weight = finalWeightAvg
	compute.Avg_Initial_Weight = initialWeightAvg
	compute.Days = days
	compute.Estimated_Weight = estimatedCurrentWeight

	daysLeft := 122 - days
	estimatedWeightLeft := float32(daysLeft) * compute.Avg_Added_Weight_Per_Day
	k := float32((estimatedWeightLeft+compute.Estimated_Weight)*100) / 100

	avgMonthlyEstimatedWeight := k / float32(4) //We based it on 5 due to the visualization of pig in UI
	daysL := 122 / 4
	avgInitialDaysOld, _ := InitialDayOldAvg()

	var i int
	var daysLe int
	for i = 1; i < 5; i++ {
		weight := avgMonthlyEstimatedWeight * float32(i)
		dayss := avgInitialDaysOld * float32(i)
		daysLe = daysLe + daysL
		fmt.Println("Weight: ", weight)

		item := Projected_Weight{
			Code:   int(dayss),
			Weight: weight,
		}
		compute.AddItem(item)

	}

	return c.JSON(compute)
}

func (weight *GeneralExpectedWeightModel) AddItem(item Projected_Weight) []Projected_Weight {
	weight.Projected_Weight = append(weight.Projected_Weight, item)
	return weight.Projected_Weight
}
