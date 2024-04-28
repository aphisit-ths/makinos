package models

import "time"

type SmsLog struct {
	Id            uint
	ApplicationId uint
	Msisdn        string
	Message       string
	Sender        string
	Type          string
	IsSuccess     bool
	CreatedAt     time.Time
}
