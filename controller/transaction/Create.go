package transaction

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils/db"
)

func Create(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	trn := model.NewTransactionModel{}
	if err := c.BodyParser(&trn); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	insertStmt := ``
	updateFnlWghtStmt := ``

	trn.Ref_ID = db.GenerateRefId(trn.FirstName, trn.LastName, len(trn.Pigs))

	var id []int

	for i, v := range trn.Pigs {

		conv, _ := strconv.Atoi(v.PigID)

		id = append(id, conv)

		updateFnlWghtStmt += fmt.Sprintf("UPDATE public.t_stock  SET final_weight = %v WHERE id = %v;", trn.Pigs[i].FinalWeight, trn.Pigs[i].PigID)

		if i == len(trn.Pigs)-1 {
			insertStmt += fmt.Sprintf("('%v',Now(), '%v', '%v', '%v', '%v', %v, %v)",
				trn.Ref_ID, trn.FirstName, trn.MiddleName, trn.LastName, trn.MobileNo, trn.PricePerKilo, v.PigID)
		} else {
			insertStmt += fmt.Sprintf("('%v',Now(), '%v','%v','%v','%v', %v, %v),",
				trn.Ref_ID, trn.FirstName, trn.MiddleName, trn.LastName, trn.MobileNo, trn.PricePerKilo, v.PigID)
		}

	}

	ids, _ := json.Marshal(id)

	sold, err1 := db.IsSold(ids, c)
	if sold {
		return err1
	}

	sqlStatement := `	
	BEGIN;
	UPDATE public.t_stock  SET status = 2 
						WHERE id = ANY (ARRAY ` + string(ids) + `);
						
	` + updateFnlWghtStmt + `					

	INSERT INTO public.t_transaction(
						ref_id, trn_date, first_name, middle_name, last_name, mobile_number, price_per_kilo, stock_id)` +
		`VALUES` + insertStmt + `;
	COMMIT;
	`
	_, err := database.CCDB.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Successfully Transacted", "data": trn})

}

func IsSold(c *fiber.Ctx) error {
	id := c.Params("id")
	type Exist struct {
		exists bool
	}

	var exist []Exist

	sqlStmnt := `SELECT EXISTS(SELECT 1 FROM public.t_transaction WHERE stock_id = ` + id + `);`
	fmt.Println(sqlStmnt)

	rows, err := database.CCDB.Query(sqlStmnt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
	}

	defer rows.Close()

	result := exist
	ex := Exist{}

	if rows.Next() {

		if err := rows.Scan(
			&ex.exists,
		); err != nil {
			return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
		}
		result = append(result, ex)

	}
	if ex.exists {
		return c.JSON(fiber.Map{"responseCode": 400, "message": "Invalid data. Please try again.", "data": nil})
	}

	// Append stock to result

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Valid Data", "data": nil})

}
