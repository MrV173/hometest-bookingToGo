package repositories

import (
	"test-gorilla-mux/models"

	"gorm.io/gorm"
)

type NationalityRepository interface {
	FindNationalities() ([]models.Nationality, error)
	GetNationality(ID int) (models.Nationality, error)
	CreateNationality(nationality models.Nationality) (models.Nationality, error)
	UpdateNationality(nationality models.Nationality) (models.Nationality, error)
	DeleteNationality(nationality models.Nationality) (models.Nationality, error)
}

func RepositoryNationality(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindNationalities() ([]models.Nationality, error) {
	var nationalities []models.Nationality
	err := r.db.Find(&nationalities).Error

	return nationalities, err
}

func (r *repository) GetNationality(ID int) (models.Nationality, error) {
	var nationality models.Nationality
	err := r.db.First(&nationality, ID).Error

	return nationality, err
}

func (r *repository) CreateNationality(nationality models.Nationality) (models.Nationality, error) {
	err := r.db.Create(&nationality).Error

	return nationality, err
}

func (r *repository) UpdateNationality(nationality models.Nationality) (models.Nationality, error) {
	err := r.db.Save(&nationality).Error

	return nationality, err
}

func (r *repository) DeleteNationality(nationality models.Nationality) (models.Nationality, error) {
	err := r.db.Delete(&nationality).Error

	return nationality, err
}
