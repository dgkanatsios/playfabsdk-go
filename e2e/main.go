package main

import (
	"fmt"
	"os"
	"strconv"

	playfab "github.com/dgkanatsios/playfabsdk-go/sdk"
	"github.com/dgkanatsios/playfabsdk-go/sdk/authentication"
	"github.com/dgkanatsios/playfabsdk-go/sdk/client"
	"github.com/dgkanatsios/playfabsdk-go/sdk/data"
	"github.com/dgkanatsios/playfabsdk-go/sdk/server"
)

const (
	testKey       string = "testCounter"
	testNumber    string = "15"
	testNumberInt int    = 15

	testStatName    string = "xp"
	testStatVersion uint32 = 0
	testStatValue   int32  = 800
)

func main() {

	titleID := os.Getenv("TitleID")
	developerSecretKey := os.Getenv("DeveloperSecretKey")

	settings := playfab.NewSettingsWithDefaultOptions(titleID)

	// CLIENT API - try to login with wrong credentials
	l := &client.LoginWithEmailAddressRequestModel{
		Email:    "wrongEmail",
		Password: "invalid",
	}
	_, err := client.LoginWithEmailAddress(settings, l)
	if err == nil {
		handleFail("Should not login with fake credentials")
	}

	// CLIENT API - try to register with inappropriate password
	l2 := &client.RegisterPlayFabUserRequestModel{
		Username: "x",
		Email:    "x",
		Password: "x",
	}
	_, err = client.RegisterPlayFabUser(settings, l2)
	if err == nil {
		handleFail("Should not register with inappropriate password")
	}

	// CLIENT API - Log in or create a user, track their PlayFabId
	l3 := &client.LoginWithCustomIDRequestModel{
		CustomId:      "GettingStartedGuide",
		CreateAccount: true,
	}

	res, err := client.LoginWithCustomID(settings, l3)
	if err != nil {
		handleFail(fmt.Sprintf("LoginWithCustomID should not return err, Error:%s", err.Error()))
	}

	playfabId := res.PlayFabId

	// CLIENT API - Test that the login call sequence sends the AdvertisingId when set
	settings.AdvertisingIdType = playfab.AdTypeAndroidId
	settings.AdvertisingIdValue = "PlayFabTestId"
	res3, err := client.LoginWithCustomID(settings, l3)
	if err != nil {
		handleFail(fmt.Sprintf("LoginWithCustomID should not return err, Error:%s", err.Error()))
	}

	// if settings.AdvertisingIdType != playfab.AD_TYPE_ANDROID_ID_Successful {
	// 	handleFail("The advertisingId was not submitted properly")
	// }

	clientSessionTicket := res3.SessionTicket

	// CLIENT API - test UpdateUserData
	l4 := &client.UpdateUserDataRequestModel{Data: map[string]string{testKey: testNumber}}
	_, err = client.UpdateUserData(settings, l4, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("UpdateUserData should not return an error, Error:%s", err.Error()))
	}

	// CLIENT API - test GetUserData
	l5 := &client.GetUserDataRequestModel{PlayFabId: playfabId}
	res5, err := client.GetUserData(settings, l5, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("GetUserData should not return an error, Error:%s", err.Error()))
	}
	val := res5.Data[testKey].Value

	if val != testNumber {
		handleFail(fmt.Sprintf("val should be equal to \"%s\", it is %s", testNumber, val))
	}

	// CLIENT API - Test Update Player Statistics
	l6 := &client.UpdatePlayerStatisticsRequestModel{
		Statistics: []client.StatisticUpdateModel{
			client.StatisticUpdateModel{
				StatisticName: testStatName,
				Version:       testStatVersion,
				Value:         testStatValue,
			},
		},
	}
	_, err = client.UpdatePlayerStatistics(settings, l6, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("UpdatePlayerStatistics should not return err, Error:%s", err.Error()))
	}

	// CLIENT API - Test Get Player Statistics
	l7 := &client.GetPlayerStatisticsRequestModel{StatisticNames: []string{testStatName}}
	res7, err := client.GetPlayerStatistics(settings, l7, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("GetPlayerStatistics should not return err, Error:%s", err.Error()))
	}
	if res7.Statistics[0].Value != testStatValue {
		handleFail(fmt.Sprintf("statistics should be equal to %d, it is %d", testStatValue, res7.Statistics[0].Value))
	}

	// CLIENT API - Get or create the given test character for the given user
	l8 := &client.ListUsersCharactersRequestModel{PlayFabId: playfabId}
	res8, err := client.GetAllUsersCharacters(settings, l8, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("GetAllUsersCharacters should not return err, Error:%s", err.Error()))
	}
	if len(res8.Characters) != 0 {
		handleFail(fmt.Sprintf("len(res8.Characters) should be equal to %d, it is %d", 0, len(res8.Characters)))
	}

	// CLIENT API - Test Get leaderboards
	l9 := &client.GetLeaderboardRequestModel{
		MaxResultsCount: 3,
		StartPosition:   0,
		StatisticName:   testStatName,
	}
	res9, err := client.GetLeaderboard(settings, l9, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("Client GetLeaderboard should not return err, Error:%s", err.Error()))
	}
	if res9.Leaderboard == nil {
		handleFail("Client Leaderboard failed")
	}
	if len(res9.Leaderboard) == 0 {
		handleFail("Client Leaderboard had insufficient entries")
	}

	// SERVER API - Test Get Leaderboards
	l10 := &server.GetLeaderboardRequestModel{
		MaxResultsCount: 3,
		StartPosition:   0,
		StatisticName:   testStatName,
	}
	res10, err := server.GetLeaderboard(settings, l10, developerSecretKey)
	if err != nil {
		handleFail(fmt.Sprintf("Server GetLeaderboard should not return err, Error:%s", err.Error()))
	}
	if res10.Leaderboard == nil {
		handleFail("Server Leaderboard failed")
	}
	if len(res10.Leaderboard) == 0 {
		handleFail("Server Leaderboard had insufficient entries")
	}

	// CLIENT API - Test Get Account Info
	l11 := &client.GetAccountInfoRequestModel{PlayFabId: playfabId}
	res11, err := client.GetAccountInfo(settings, l11, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("GetAccountInfo should not return err, Error:%s", err.Error()))
	}
	if res11.AccountInfo.TitleInfo.Origination == "" {
		handleFail("Origination should not be empty string")
	}

	// CLIENT API - Test Cloud Script
	l12 := &client.ExecuteCloudScriptRequestModel{FunctionName: "helloWorld", FunctionParameter: "1234"}
	res12, err := client.ExecuteCloudScript(settings, l12, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("ExecuteCloudScript should not return err, Error:%s", err.Error()))
	}
	if res12.FunctionResult == nil {
		handleFail("Function should return a value")
	}

	// CLIENT API - Test Cloud Script
	l13 := &client.ExecuteCloudScriptRequestModel{FunctionName: "throwError"}
	res13, err := client.ExecuteCloudScript(settings, l13, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("ExecuteCloudScript with 'throwError' should not return err, Error:%s", err.Error()))
	}
	if res13.FunctionResult != nil {
		handleFail("Function should have failed")
	}

	// CLIENT API - Test  that the client can publish custom PlayStream events
	l14 := &client.WriteClientPlayerEventRequestModel{
		EventName: "ForumPostEvent",
		Body: map[string]interface{}{
			"Subject": "My First Post",
			"Body":    "This is my awesome post",
		},
	}
	_, err = client.WritePlayerEvent(settings, l14, clientSessionTicket)
	if err != nil {
		handleFail(fmt.Sprintf("WritePlayerEvent should not return err, Error:%s", err.Error()))
	}

	// ENTITY API - Verify that a client login can be converted into an entity token
	l15 := &authentication.GetEntityTokenRequestModel{}
	res15, err := authentication.GetEntityToken(settings, l15, "", clientSessionTicket, "")
	if err != nil {
		handleFail(fmt.Sprintf("GetEntityToken should not return err, Error:%s", err.Error()))
	}
	if res15.Entity.Id == "" {
		handleFail("entityId should be defined")
	}
	if res15.Entity.Type == "" {
		handleFail("entityType should be defined")
	}
	entityToken := res15.EntityToken
	entity := res15.Entity

	// ENTITY API - Test a sequence of calls that modifies entity objects
	l16 := &data.SetObjectsRequestModel{Entity: &data.EntityKeyModel{
		Id:   entity.Id,
		Type: entity.Type,
	}, Objects: []data.SetObjectModel{
		data.SetObjectModel{
			ObjectName: testKey,
			DataObject: testNumberInt,
		},
	}}
	_, err = data.SetObjects(settings, l16, entityToken)
	if err != nil {
		handleFail(fmt.Sprintf("SetObjects should not return err, Error:%s", err.Error()))
	}

	l17 := &data.GetObjectsRequestModel{Entity: &data.EntityKeyModel{
		Id:   entity.Id,
		Type: entity.Type,
	}, EscapeObject: true}
	res17, err := data.GetObjects(settings, l17, entityToken)
	if err != nil {
		handleFail(fmt.Sprintf("GetObjects should not return err, Error:%s", err.Error()))
	}
	obj, ok := res17.Objects[testKey]
	if !ok {
		handleFail("Object testKey should be included")
	}
	valInt, errConv := strconv.Atoi(obj.EscapedDataObject)
	if errConv != nil {
		handleFail("Should be able to convert EscapedDataObject to int")
	}

	if valInt != testNumberInt {
		handleFail(fmt.Sprintf("val should be %d but it is %d", testNumberInt, valInt))
	}
}

func handleFail(msg string) {
	panic(msg)
}
