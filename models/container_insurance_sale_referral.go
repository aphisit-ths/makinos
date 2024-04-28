package models

import "time"

type ContainerInsuranceSaleReferral struct {
	ID                       int64     `json:"id"`
	PhoneNumber              string    `json:"phone_number"`
	ContainerInsuranceSaleID int64     `json:"container_insurance_sale_id"`
	CreatedAt                time.Time `json:"created_at"`
	CreatedBy                string    `json:"created_by"`
	UpdatedAt                time.Time `json:"updated_at"`
	UpdatedBy                string    `json:"updated_by"`
	IsDeleted                bool      `json:"is_deleted"`
}

func (ContainerInsuranceSaleReferral) TableName() string {
	return "insurance.container_insurance_sale_referral"
}
