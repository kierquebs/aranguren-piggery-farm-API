package appointment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
	"github.com/kierquebs/aranguren-piggery-farm-API/utils"
)

func Update(c *fiber.Ctx) error {
	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")
	updateApt := new(model.UpdateAppointmentModel)
	utils.BodyParser(c, updateApt)

	sqlStatement := `UPDATE public.t_appointment
	SET mobile_number=$1, first_name=$2, middle_name=$3, last_name=$4, email_address=$5, appointment_date=$6, status=$7
	WHERE id=$8;`
	_, err := database.CCDB.Exec(sqlStatement, updateApt.Mobile_Number, updateApt.First_Name, updateApt.Middle_Name, updateApt.Last_Name, updateApt.Appointment_Date, updateApt.Status)
	if err != nil {
		return c.JSON(fiber.Map{"responseCode": 500, "message": "Error updating appointment.", "error": err.Error()})
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "Appointment successfully updated."})

}
