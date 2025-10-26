package calculationservice

import "gorm.io/gorm"

type CalculationRepository interface {
	CreateCalculation(calc Calculation) error
	GetAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(calc Calculation) error
	DeleteCalculation(id string) error
}

type calcRespository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRespository{db: db}
}

func (r *calcRespository) CreateCalculation(calc Calculation) error {
	return r.db.Create(&calc).Error
}

func (r *calcRespository) GetAllCalculations() ([]Calculation, error) {
	var calculations []Calculation

	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *calcRespository) GetCalculationByID(id string) (Calculation, error) {
	var calc Calculation

	err := r.db.First(&calc, "id = ?", id).Error
	return calc, err
}

func (r *calcRespository) UpdateCalculation(calc Calculation) error {
	return r.db.Save(&calc).Error
}

func (r *calcRespository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculation{}, "id = ?", id).Error
}
