package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name       string  `json:"name"`
		Country    string  `json:"country"`
		Local_time string  `json:"localtime"`
		Latitude   float64 `json:"lat"`
		Longitude  float64 `json:"lon"`
	} `json:"location"`

	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func colorCondition(text string) { //Just some of them, there is a csv for the same
	switch text {
	case "Sunny":
		color.Yellow("It is %s Outside", text)
	case "Rainy":
		color.Blue("It is %s Outside", text)
		//return color.New(color.FgBlue).Sprint(text)
	case "Cloudy":
		color.White("It is %s Outside", text)
	default:
		color.Cyan("It is %s Outside", text)
	}
}

func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEY not found in .env")
		return
	}

	q := "Mumbai" //default
	if len(os.Args) < 2 {
		fmt.Println("Error: No argument provided.")
		fmt.Println("Usage: go run main.go <location>")
		return
	} else {
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + API_KEY + "&q=" + q + "&aqi=no")

	if err != nil { //Error exists (Non Null)
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body) //read the data (got)
	if err != nil {
		panic(err)
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}
	location, current := weather.Location, weather.Current
	fmt.Printf("%s %s %s %0.3f %0.3f %0.1f",
		location.Name,
		location.Country,
		location.Local_time,
		location.Latitude,
		location.Longitude,
		current.TempC,
	)
	fmt.Println()
	colorCondition(current.Condition.Text)
}
