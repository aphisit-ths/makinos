package repositories

import (
	"makino/models"
	"time"

	"github.com/jaswdr/faker"
)

type ReferralRepositoryMock struct {
}

func NewReferralRepositoryMock() ReferralRepository {
	return &ReferralRepositoryMock{}
}

func (r *ReferralRepositoryMock) GetSaleReferralByDate(datetime string) ([]models.ContainerInsuranceSaleReferral, error) {
	var referrals []models.ContainerInsuranceSaleReferral
	for i := 0; i < 100; i++ {
		fake := faker.New()
		phoneNumber := fake.Phone().Number()
		var ref = models.ContainerInsuranceSaleReferral{
			ID:                       fake.Int64(),
			PhoneNumber:              phoneNumber,
			ContainerInsuranceSaleID: fake.Int64(),
			CreatedAt:                time.Now(),
			CreatedBy:                "1",
			UpdatedAt:                time.Now(),
			UpdatedBy:                "1",
			IsDeleted:                false,
		}
		referrals = append(referrals, ref)
	}
	return referrals, nil
}
