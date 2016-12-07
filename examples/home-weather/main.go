package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/kyokomi/emoji"
)

var jsonWeather = `{
"System" : "HomeWeather",
"Rooms" : {
   "Kitchen": "27",
   "Bedroom": "26",
   "Living": "27",
   "Tarrace": "18"}
}`

type RoomType int

const (
	KITCHEN RoomType = 1 << iota
	BEDROOM
	LIVING
	TARRACE
)

func getHomeWeatherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, jsonWeather)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "---- Home Weather System ----\n")
	tempKitchen := getRoomTemperature(KITCHEN)
	tempBedroom := getRoomTemperature(BEDROOM)
	tempLiving := getRoomTemperature(LIVING)
	tempTarrace := getRoomTemperature(TARRACE)

	fmt.Fprintf(w, "Kitchen : %d %s\n", tempKitchen, getEmoji(tempKitchen))
	fmt.Fprintf(w, "Bedroom : %d %s\n", tempBedroom, getEmoji(tempBedroom))
	fmt.Fprintf(w, "Living  : %d %s\n", tempLiving, getEmoji(tempLiving))
	fmt.Fprintf(w, "Tarrace : %d %s\n", tempTarrace, getEmoji(tempTarrace))
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func getEmoji(temp int) (emojiString string) {
	switch temp {
	case 7, 8, 9, 10:
		emojiString = emoji.Sprint(" :snowflake: ")
	case 11, 12, 13, 14, 15:
		emojiString = emoji.Sprint(" :cloud_with_snow: ")
	case 16, 17, 18, 19, 20:
		emojiString = emoji.Sprint(" :cool: ")
	case 21, 22, 23, 24, 25:
		emojiString = emoji.Sprint(" :cloud: ")
	case 26, 27, 28, 29:
		emojiString = emoji.Sprint(" :sunny: ")
	case 30, 31, 32:
		emojiString = emoji.Sprint(" :fire: ")
	}
	return emojiString
}

func getRoomTemperature(room RoomType) (temp int) {
	switch room {
	case KITCHEN:
		temp = randInt(20, 32)
	case BEDROOM:
		temp = randInt(7, 32)
	case LIVING:
		temp = randInt(7, 32)
	case TARRACE:
		temp = randInt(7, 24)
	}
	return temp
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	http.HandleFunc("/getHomeWeather", getHomeWeatherHandler)
	fmt.Println("Server is listening at : 8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
