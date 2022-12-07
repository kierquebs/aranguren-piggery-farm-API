package model

type CreateAppointmentModel struct {
	Mobile_Number    string `json:"mobile_number"`
	First_Name       string `json:"first_name"`
	Middle_Name      string `json:"middle_name"`
	Last_Name        string `json:"last_name"`
	Email_Address    string `json:"email_address"`
	Appointment_Date string `json:"appointment_date"`
	Status           int    `json:"status"`
}
