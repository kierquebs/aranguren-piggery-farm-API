package transaction

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/kierquebs/aranguren-piggery-farm-API/database"

	"github.com/gofiber/fiber/v2"
)

type Stock struct {
	ID             int32   `json:"id"`
	Added_Date     string  `json:"added_date"`
	Qr_Code        string  `json:"qr_code"`
	Final_Weight   float64 `json:"final_weight"`
	Initial_Weight int     `json:"initial_weight"`
}

type StockSlice []Stock

type ViewTransactionModel struct {
	RefID        string     `json:"ref_id"`
	Trn_Date     string     `json:"trn_date"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	MiddleName   string     `json:"middle_name"`
	PricePerKilo float64    `json:"price_per_kilo"`
	MobileNo     string     `json:"mobile_no"`
	Stocks       StockSlice `json:"stock"`
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (s Stock) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (s *StockSlice) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion failed")
}

var trnModel []ViewTransactionModel

func Find(c *fiber.Ctx) error {

	trnRows, err := database.CCDB.Query(`	SELECT  
												t.ref_id, 
												t.trn_date, 
												t.first_name, 
												t.middle_name, 
												t.last_name, 
												t.mobile_number, 
												t.price_per_kilo, 
												jsonb_agg(json_build_object(
													'id',s.id,
													'added_date',s.added_date,
													'qr_code',s.qr_code,
													'final_weight',s.final_weight,
													'initial_weight',s.initial_weight
												)) as stock
											FROM public.t_transaction t
											JOIN public.t_stock s ON s.id = t.stock_id
											GROUP BY
												t.ref_id, 
												t.trn_date, 
												t.first_name, 
												t.middle_name, 
												t.last_name, 
												t.mobile_number, 
												t.price_per_kilo
									`)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
	}

	defer trnRows.Close()

	result := trnModel

	for trnRows.Next() {
		trn := ViewTransactionModel{}
		if err := trnRows.Scan(
			&trn.RefID,
			&trn.Trn_Date,
			&trn.FirstName,
			&trn.MiddleName,
			&trn.LastName,
			&trn.MobileNo,
			&trn.PricePerKilo,
			&trn.Stocks,
		); err != nil {
			return err // Exit if we get an error
		}

		result = append(result, trn)

	}

	// Return Stock in JSON format
	return c.JSON(fiber.Map{"responseCode": 200, "message": "Details fetched succesfully.", "data": result})

}
