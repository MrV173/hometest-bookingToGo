package repositories

import (
	"test-gorilla-mux/models"

	"gorm.io/gorm"
)

type FamilyRepository interface {
	FindFamilies() ([]models.Family, error)
	GetFamily(ID int) (models.Family, error)
	CreateFamily(family models.Family) (models.Family, error)
	UpdateFamily(family models.Family) (models.Family, error)
	DeleteFamily(family models.Family) (models.Family, error)
}

func RepositoryFamily(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFamilies() ([]models.Family, error) {
	var families []models.Family
	err := r.db.Find(&families).Error

	return families, err
}

func (r *repository) GetFamily(ID int) (models.Family, error) {
	var family models.Family
	err := r.db.First(&family).Error

	return family, err
}

func (r *repository) CreateFamily(family models.Family) (models.Family, error) {
	err := r.db.Create(&family).Error

	return family, err
}

func (r *repository) UpdateFamily(family models.Family) (models.Family, error) {
	err := r.db.Save(&family).Error

	return family, err
}

func (r *repository) DeleteFamily(family models.Family) (models.Family, error) {
	err := r.db.Delete(&family).Error

	return family, err
}
