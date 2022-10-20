package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
)

func GenerateRefId(fname, lname string, quantity int) string {
	fInitial := fname[0:1]
	lInitial := lname[0:1]
	return fmt.Sprintf("%v-%v%v-%v", CurrentTime(), strings.ToUpper(fInitial), strings.ToUpper(lInitial), quantity)

}

// CurrentTime
func CurrentTime() string {
	// with the help of time.Now()
	// store the local time
	t := time.Now()

	// print location and local time
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
	}
	trimmedSpace := strings.Replace(t.In(location).Format("2006-01-02 15:04:05"), " ", "", -1)
	trimmedDash := strings.Replace(trimmedSpace, "-", "", -1)
	trimmedColon := strings.Replace(trimmedDash, ":", "", -1)

	return trimmedColon
}

func IsSold(ids []byte, c *fiber.Ctx) (bool, error) {

	type Exist struct {
		exists bool
	}

	var exist []Exist

	sqlStmnt := `SELECT EXISTS(SELECT 1 FROM public.t_transaction WHERE stock_id = ANY(ARRAY` + string(ids) + `));`
	fmt.Println(sqlStmnt)

	rows, err := database.CCDB.Query(sqlStmnt)
	if err != nil {
		return true, c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
	}

	defer rows.Close()

	result := exist
	ex := Exist{}

	if rows.Next() {

		if err := rows.Scan(
			&ex.exists,
		); err != nil {
			return true, c.Status(500).JSON(fiber.Map{"responseCode": 500, "message": "Error: " + err.Error(), "data": nil})
		}
		result = append(result, ex)

	}
	if ex.exists {
		return true, c.JSON(fiber.Map{"responseCode": 400, "message": "Invalid data. Please try again.", "data": nil})
	}

	// Append stock to result

	return false, c.JSON(fiber.Map{"responseCode": 200, "message": "Valid Data", "data": nil})

}
