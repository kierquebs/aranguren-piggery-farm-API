package model

type CreateStockModel struct {
	Added_By        int32   `json:"added_by"`
	Initial_Weight  float32 `json:"initial_weight"`
	Initial_Day_Old int32   `json:"initial_day_old"`
}

//ViewStockModel  this model is used by the Database public.view.view_t_stock
type ViewStockModel struct {
	ID                       int32   `json:"id"`
	Added_Date               string  `json:"added_date"`
	Last_Update_Date         *string `json:"last_updated_date"`
	Initial_Weight           float32 `json:"initial_weight"`
	Estimated_Current_Weight float32 `json:"estimated_current_weight"`
	Initial_Day_Old          int32   `json:"-"`
	Current_DateTime         string  `json:"-"`
	Age_By_Days              int     `json:"age_by_days"`
	Description              string  `json:"description"`
	Status_Description       string  `json:"status_description"`
	Status                   int32   `json:"-"`
}
