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

	db.CurrentTime()
	insertStmt := ``
	refId := db.GenerateRefId(trn.FirstName, trn.LastName, len(trn.Pigs))
	var id []int

	for i, v := range trn.Pigs {

		conv, _ := strconv.Atoi(v.PigID)

		id = append(id, conv)

		if i == len(trn.Pigs)-1 {
			insertStmt += fmt.Sprintf("('%v',Now(), '%v', '%v', '%v', '%v', %v, %v)",
				refId, trn.FirstName, trn.MiddleName, trn.LastName, trn.MobileNo, trn.PricePerKilo, v.PigID)
		} else {
			insertStmt += fmt.Sprintf("('%v',Now(), '%v','%v','%v','%v', %v, %v),",
				refId, trn.FirstName, trn.MiddleName, trn.LastName, trn.MobileNo, trn.PricePerKilo, v.PigID)
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

	return c.JSON(fiber.Map{"responseCode": 200, "message": sqlStatement})

}
