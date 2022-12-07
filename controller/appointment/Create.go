package appointment

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kierquebs/aranguren-piggery-farm-API/database"
	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

func Create(c *fiber.Ctx) error {

	c.Set(fiber.HeaderAccessControlAllowOrigin, "*")

	newAppointment := model.CreateAppointmentModel{}
	if err := c.BodyParser(&newAppointment); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}

	_, err := database.CCDB.Exec(`INSERT INTO public.t_appointment(
		mobile_number,
		first_name,
		middle_name,
		last_name,
		email_address,
		date_added,
		appointment_date,
		status)
		VALUES ($1,$2,$3,$4,$5,Now() AT TIME ZONE 'Asia/Manila',$6,1);`,
		newAppointment.Mobile_Number,
		newAppointment.First_Name,
		newAppointment.Middle_Name,
		newAppointment.Last_Name,
		newAppointment.Email_Address,
		newAppointment.Appointment_Date,
	)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
		return err
	}

	return c.JSON(fiber.Map{"responseCode": 200, "message": "New appointment successfully added."})

}
