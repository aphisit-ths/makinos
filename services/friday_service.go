package services

import (
	"fmt"
	"makino/businesslogic"
	"makino/models"
)

type IFridayService interface {
	SendSmS(name string) error
}

type FridayService struct {
	setting models.AppSettings
}

func NewFridayService(setting models.AppSettings) FridayService {
	return FridayService{setting: setting}
}

func (FridayService) SendSmS(name string) {
	fmt.Println(businesslogic.CreateSmsWordingWithoutCustomer(name))
}
