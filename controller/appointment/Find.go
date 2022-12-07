package appointment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

var vapt []model.ViewAppointmentModel

func ListAll(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	rows, err := database.CCDB.Query(`SELECT 
										id, 
										mobile_number, 
										first_name, 
										middle_name, 
										last_name, 
										email_address, 
										date_added, 
										appointment_date, 
										status
										FROM public.t_appointment;
									`)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	result := vapt

	for rows.Next() {
		appointment := model.ViewAppointmentModel{}
		if err := rows.Scan(
			&appointment.ID,
			&appointment.Mobile_Number,
			&appointment.First_Name,
			&appointment.Middle_Name,
			&appointment.Last_Name,
			&appointment.Email_Address,
			&appointment.Date_Added,
			&appointment.Appointment_Date,
			&appointment.Status,
		); err != nil {
			return err // Exit if we get an error
		}

		// Append stock to result
		result = append(result, appointment)
	}
	// Return Stock in JSON format
	return c.JSON(result)
}
