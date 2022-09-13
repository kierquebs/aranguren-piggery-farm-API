package model

type ImageUpload struct {
	Data struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		URLViewer  string `json:"url_viewer"`
		URL        string `json:"url"`
		DisplayURL string `json:"display_url"`
		Width      string `json:"width"`
		Height     string `json:"height"`
		Size       int    `json:"size"`
		Time       string `json:"time"`
		Expiration string `json:"expiration"`
		Image      struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"image"`
		Thumb struct {
			Filename  string `json:"filename"`
			Name      string `json:"name"`
			Mime      string `json:"mime"`
			Extension string `json:"extension"`
			URL       string `json:"url"`
		} `json:"thumb"`
		DeleteURL string `json:"delete_url"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
