package model

type StockModel struct {
	InitialWeight int32 `json:"initial_weight"`
	CurrentWeight int32 `json:"current_weight"`
	Category      Category
	AddedBy       AddedBy
	Price         Price
}

type Category struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AddedBy struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type Price struct {
	Amount          string `json:"amount"`
	LastUpdatedBy   LastUpdatedBy
	LastUpdatedDate string `json:"last_updated_date"`
}

type LastUpdatedBy struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}
