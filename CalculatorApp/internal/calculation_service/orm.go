package calculationservice

type Calculation struct {
	ID         string `gorm:"PrimaryKey" json:"id"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

type CalculationRequest struct {
	Expression string `json:"expression"`
}
