package models

type Application struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Date     string `json:"date"`
	Link     string `json:"link"`
	Notes    string `json:"notes"`
	Status   string `json:"status"`
}
