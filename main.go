package main

import (
	"makino/businessflow"
	"makino/config"
	"makino/repositories"
	"time"
)

func main() {
	setting := config.GetSettings()

	_, ctx, client, cancel := config.ConnectToCustomerDB(setting)

	defer client.Disconnect(ctx)
	defer cancel()

	postgresDb := config.ConnectToPostgresDB(setting)

	instance, err := postgresDb.DB()
	if err != nil {
		panic(err)
	}

	defer instance.Close()

	//personRepository := repositories.NewPersonInfoRepository(&ctx, customerDB)
	//referralRepository := repositories.NewReferralRepository(postgresDb)
	referralRepoMock := repositories.NewReferralRepositoryMock()
	businessFlow := businessflow.NewReferralBusinessFlow(referralRepoMock)

	businessFlow.SendSMSToClient(time.DateTime)
}
