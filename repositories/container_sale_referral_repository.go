package repositories

import (
	"errors"
	"makino/models"

	"gorm.io/gorm"
)

type ReferralRepository interface {
	GetSaleReferralByDate(date string) ([]models.ContainerInsuranceSaleReferral, error)
}

type referralRepositoryImpl struct {
	database *gorm.DB
}

func NewReferralRepository(db *gorm.DB) ReferralRepository {
	return &referralRepositoryImpl{database: db}
}

func (r *referralRepositoryImpl) GetSaleReferralByDate(date string) ([]models.ContainerInsuranceSaleReferral, error) {
	database := r.database
	var referrals []models.ContainerInsuranceSaleReferral
	if err := database.Find(&referrals, "created_at <= ?", date).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, nil
		}
		return nil, err
	}
	return referrals, nil
}
