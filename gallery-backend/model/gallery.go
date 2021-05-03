package model

type Paint struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Exhibition struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Paints      []Paint `json:"paints"`
}
