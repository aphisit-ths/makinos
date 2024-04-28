package businesslogic

import (
	"fmt"
	"os"
)

func GetAppEnvironment() string {
	var response string
	if os.Getenv("GOLANG_ENVIRONMENT") != "" {
		response = os.Getenv("GOLANG_ENVIRONMENT")
	} else {
		response = "Local"
	}
	fmt.Println("### batch working on", response)

	return response
}
