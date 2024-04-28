package businessflow

import (
	"fmt"
	"makino/repositories"
)

type ReferralBusinessFlow struct {
	referralRepository repositories.ReferralRepository
}

func NewReferralBusinessFlow(referralRepository repositories.ReferralRepository) *ReferralBusinessFlow {
	return &ReferralBusinessFlow{
		referralRepository: referralRepository,
	}
}

func (r *ReferralBusinessFlow) SendSMSToClient(date string) {
	fmt.Println("### Start batch process ")
	referrals, err := r.referralRepository.GetSaleReferralByDate(date)

	if err != nil {
		fmt.Println("### Something went wrong ", err)
		panic(err)
	}
	if referrals == nil {
		fmt.Println("### Have friend get friend for today !")
		return
	}

}
