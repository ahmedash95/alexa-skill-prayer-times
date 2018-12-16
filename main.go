package main

import (
	"fmt"
	"strings"

	"github.com/ahmedash95/alexa-prayer-times/prayer"
	alexa "github.com/ahmedash95/amazon-alexa-sdk"
)

func main() {
	alexa := alexa.New()
	alexa.AddIntentResponse("ListPrayingTimes", handleListPrayingTimesResponse)
	alexa.AddIntentResponse("SetUserInfo", handleSetUserInfoIntent)
	alexa.Run("3000")
}

func handleListPrayingTimesResponse(request alexa.Request) alexa.StringPayload {
	var payload alexa.StringPayload
	user := alexa.GetUser()
	if user.Name == "" {
		return handleAskForUserInfo(request)
	}

	city := user.Location
	country, _ := GetCountryByCityName(user.Location)
	list := prayer.GetList(city, country)

	removeCET := func(s string) string {
		return strings.Replace(s, " (CET)", "", 1)
	}

	text := "Here is today praying times"
	text += fmt.Sprintf("\nElFajr: %s", removeCET(list.Timings.Fajr))
	text += fmt.Sprintf("\nSunrise: %s", removeCET(list.Timings.Sunrise))
	text += fmt.Sprintf("\nElDoohr: %s", removeCET(list.Timings.Dhuhr))
	text += fmt.Sprintf("\nElAsr: %s", removeCET(list.Timings.Asr))
	text += fmt.Sprintf("\nElMaghrib: %s", removeCET(list.Timings.Maghrib))
	text += fmt.Sprintf("\nElIsha: %s", removeCET(list.Timings.Isha))

	payload.Title = "praying times list"
	payload.Text = text

	return payload
}

func handleAskForUserInfo(request alexa.Request) alexa.StringPayload {
	return alexa.StringPayload{
		Title: "ask for user name",
		Text:  "What's your name and location ?",
	}
}

func handleSetUserInfoIntent(request alexa.Request) alexa.StringPayload {
	slots := request.Body.Intent.Slots

	name := slots["name"]
	location := slots["location"]

	alexa.SetUser(alexa.User{
		Name:     name.Value,
		Location: location.Value,
	})

	var payload alexa.StringPayload
	payload.Title = "set user info"
	payload.Text = fmt.Sprintf("Hello %s, I have saved your info", name.Value)
	return payload
}
