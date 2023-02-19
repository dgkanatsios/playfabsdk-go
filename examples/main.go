package main

import (
	"fmt"

	playfab "github.com/dgkanatsios/playfabsdk-go/sdk"
	client "github.com/dgkanatsios/playfabsdk-go/sdk/client"
)

func main() {
	settings := playfab.NewSettingsWithDefaultOptions("YOUR_TITLEID_HERE")

	loginData := &client.LoginWithCustomIDRequestModel{
		CustomId:      "GettingStartedGuide",
		CreateAccount: true,
	}
	res, err := client.LoginWithCustomID(settings, loginData)

	if err != nil {
		fmt.Printf("Login Error: %s\n", err.Error())
		return
	}

	fmt.Printf("Login OK, SessionTicket: %s\n", res.SessionTicket)

	sessionTicket := res.SessionTicket

	updateStatisticsData := &client.UpdatePlayerStatisticsRequestModel{
		Statistics: []client.StatisticUpdateModel{
			client.StatisticUpdateModel{
				StatisticName: "xp",
				Version:       1,
				Value:         800,
			},
		},
	}

	res2, err := client.UpdatePlayerStatistics(settings, updateStatisticsData, sessionTicket)

	if err != nil {
		fmt.Printf("Update Statistics Error: %s \n", err.Error())
		return
	}

	fmt.Printf("Update Statistics OK: %#v\n", res2)
}
