package model

type CreateUserModel struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	First_Name  string `json:"first_name"`
	Middle_Name string `json:"middle_name"`
	Last_Name   string `json:"last_name"`
}

type LoginUserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
