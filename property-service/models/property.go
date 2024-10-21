package models

type Property struct {
	ID         uint              `json:"id" gorm:"primaryKey"`
	ContractID uint              `json:"contract_id"`
	Type       string            `json:"type"`
	Location   string            `json:"location"`
	Specs      map[string]string `json:"specs"`
}
