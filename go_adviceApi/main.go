package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Advice struct {
	Slip struct {
		Id  int64  `json:"id"`
		Adv string `json:"advice"`
	} `json:"slip"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("The API Endpoint could not be retrieved!")
		return
	}

	API_EP := os.Getenv("API_URL")
	if API_EP == "" {
		fmt.Println("Invalid / Could Not Find API Endpoint")
	}

	res, err := http.Get(API_EP)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() //Free Resources

	if res.StatusCode != 200 {
		panic("Advice API not available")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var advice Advice
	err = json.Unmarshal(body, &advice)

	if err != nil {
		panic(err)
	}
	adv := advice.Slip.Adv

	fmt.Printf("The Advice for the Day is %s", adv)

}
