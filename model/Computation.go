package model

type GeneralExpectedWeightModel struct {
	Avg_Added_Date           string  `json:"avg_added_date"`
	Avg_Initial_Weight       float32 `json:"avg_initial_weight"`
	Avg_Final_Weight         float32 `json:"avg_final_weight"`
	Avg_Added_Weight_Per_Day float32 `json:"avg_added_weight_per_day"`
	Days                     int     `json:"days"`
	Estimated_Weight         float32 `json:"estimated_weight"`
	Projected_Weight         []Projected_Weight
}

type Projected_Weight struct {
	Code   int     `json:"code"`
	Weight float32 `json:"weight"`
}
