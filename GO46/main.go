package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
	HTTP Response
	{
	"activity": "Learn Express.js",
	"availability": 0.25,
	"type": "education",
	"participants": 1,
	"price": 0.1,
	"accessibility": "Few to no challenges",
	"duration": "hours",
	"kidFriendly": true,
	"link": "https://expressjs.com/",
	"key": "3943506"
	}
*/

type boredResponse struct {
	Activity      string  `json:"activity"`
	Availability  float64 `json:"availability"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Accessibility string  `json:"accessibility"`
	Duration      string  `json:"duration"`
	KidFriendly   bool    `json:"kidFriendly"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
}

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://bored-api.appbrewery.com/random", nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing request:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error invalid status code:", res.Status)
		return
	}

	var boredRes boredResponse
	if err := json.NewDecoder(res.Body).Decode(&boredRes); err != nil {
		fmt.Println("error decoding response:", err)
		return
	}
	fmt.Printf("%+v\n", boredRes)

}
