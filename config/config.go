package config

import (
	"encoding/json"
	"fmt"
	"makino/businesslogic"
	"makino/models"
	"os"
)

func GetSettings() models.AppSettings {
	printArt()
	configName := fmt.Sprintf("appsettings.%s.json", businesslogic.GetAppEnvironment())
	file, err := os.ReadFile(configName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var setting models.AppSettings
	err = json.Unmarshal(file, &setting)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("### Reading", configName, " success")
	return setting
}

func printArt() {
	art :=
		`
		/*  ___      ___       __       __   ___   __    _____  ___      ______    */
		/* |"  \    /"  |     /""\     |/"| /  ") |" \  (\"   \|"  \    /    " \   */
		/*  \   \  //   |    /    \    (: |/   /  ||  | |.\\   \    |  // ____  \  */
		/*  /\\  \/.    |   /' /\  \   |    __/   |:  | |: \.   \\  | /  /    ) :) */
		/* |: \.        |  //  __'  \  (// _  \   |.  | |.  \    \. |(: (____/ //  */
		/* |.  \    /:  | /   /  \\  \ |: | \  \  /\  |\|    \    \ | \        /   */
		/* |___|\__/|___|(___/    \___)(__|  \__)(__\_|_)\___|\____\)  \"_____/    */
		`

	fmt.Println(art)
}
