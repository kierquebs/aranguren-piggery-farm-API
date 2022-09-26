package model

type CreateStockModel struct {
	Added_By        int32   `json:"added_by"`
	Initial_Weight  float32 `json:"initial_weight"`
	Initial_Day_Old int32   `json:"initial_day_old"`
}

//ViewStockModel  this model is used by the Database public.view.view_t_stock
type ViewStockModel struct {
	ID                              int32   `json:"id"`
	Added_Date                      string  `json:"added_date"`
	Added_By                        string  `json:"added_by"`
	Last_Update_Date                *string `json:"last_updated_date"`
	Updated_By                      *string `json:"updated_by"`
	Initial_Weight                  float32 `json:"initial_weight"`
	Current_Weight                  float32 `json:"current_weight"`
	Type                            string  `json:"type"`
	Type_Description                string  `json:"type_description"`
	Current_Price                   float32 `json:"current_price"`
	Current_Price_Last_Updated_Date *string `json:"current_price_last_updated_date"`
	Remarks                         *string `json:"remarks"`
}
