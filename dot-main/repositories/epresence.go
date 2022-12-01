package repositories

import (
	"journey/models"

	"gorm.io/gorm"
)

type EpresenceRepository interface {
	FindEpresences() ([]models.Epresence, error)
	FindEpresencesbyUserId(User_id int) ([]models.Epresence, error)
	GetEpresence(ID int) (models.Epresence, error)
	CreateEpresence(Epresence models.Epresence) (models.Epresence, error)
	UpdateEpresence(Epresence models.Epresence) (models.Epresence, error)
	DeleteEpresence(Epresence models.Epresence) (models.Epresence, error)
}

func RepositoryArtikel(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpresences() ([]models.Epresence, error) {
	var epresences []models.Epresence
	err := r.db.Preload("User").Find(&epresences).Error

	return epresences, err
}
func (r *repository) FindEpresencesbyUserId(User_id int) ([]models.Epresence, error) {
	var eprsence []models.Epresence
	err := r.db.Preload("User").Find(&eprsence, "user_id= ?", User_id).Error

	return eprsence, err
}

func (r *repository) GetEpresence(ID int) (models.Epresence, error) {
	var eprsence models.Epresence
	err := r.db.Preload("User").First(&eprsence, ID).Error

	return eprsence, err

}

func (r *repository) CreateEpresence(Epresence models.Epresence) (models.Epresence, error) {
	err := r.db.Preload("User").Create(&Epresence).Error

	return Epresence, err
}

func (r *repository) UpdateEpresence(Epresence models.Epresence) (models.Epresence, error) {
	err := r.db.Save(&Epresence).Error

	return Epresence, err
}

func (r *repository) DeleteEpresence(Epresence models.Epresence) (models.Epresence, error) {
	err := r.db.Delete(&Epresence).Error

	return Epresence, err
}
