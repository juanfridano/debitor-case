package models

type Person struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	ContractIDs []uint `json:"contract_ids"`
}
