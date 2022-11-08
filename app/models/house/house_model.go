package house

import (
	"Spider/app/models"
	"Spider/pkg/database"
)

type House struct {
	models.BaseModel

	City       string  `json:"city,omitempty"`
	Name       string  `json:"name,omitempty"`
	TotalPrice float64 `json:"total_price,omitempty"`
	Address    string  `json:"address,omitempty"`
	DanPrice   string  `json:"dan_price,omitempty"`
	Info       string  `json:"info,omitempty"`
	Url        string  `json:"url,omitempty"`
	Area       string  `json:"area,omitempty"`

	models.CommonTimestampsField
}

func (H *House) Create() {
	database.DB.Create(&H)
}
