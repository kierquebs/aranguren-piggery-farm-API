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

type ViewTransactionModel struct {
	RefID        string  `json:"ref_id"`
	Trn_Date     string  `json:"trn_date"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	MiddleName   string  `json:"middle_name"`
	PricePerKilo float64 `json:"price_per_kilo"`
	MobileNo     string  `json:"mobile_no"`
	Pigs         struct {
		ID             string  `json:"id"`
		Added_Date     string  `json:"added_date"`
		Qr_Code        string  `json:"qr_code"`
		Final_Weight   float64 `json:"final_weight"`
		Initial_Weight int     `json:"initial_weight"`
	} `json:"pigs"`
}
