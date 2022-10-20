package model

type NewTransactionModel struct {
	Pigs []struct {
		PigID         string  `json:"pig_id"`
		QrCode        string  `json:"qr_code"`
		FinalWeight   float64 `json:"final_weight"`
		InitialWeight float64 `json:"initial_weight"`
	} `json:"pigs"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MiddleName   string  `json:"middle_name"`
	PricePerKilo float64 `json:"price_per_kilo"`
	MobileNo     string  `json:"mobile_no"`
}
