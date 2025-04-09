package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type StandingsResponse struct {
	MRData struct {
		StandingsTable struct {
			StandingsLists []struct {
				DriverStandings []struct {
					Position string `json:"position"`
					Points   string `json:"points"`
					Driver   struct {
						FamilyName string `json:"familyName"`
					} `json:"Driver"`
				} `json:"DriverStandings"`
			} `json:"StandingsLists"`
		} `json:"StandingsTable"`
	} `json:"MRData"`
}

func main() {
	url := "https://api.jolpi.ca/ergast/f1/2025/driverstandings/?format=json"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))
	var standings StandingsResponse
	if err := json.Unmarshal(body, &standings); err != nil {
		panic(err)
	}

	fmt.Printf("%-20s %s\n", "Driver", "Points")
	for _, standing := range standings.MRData.StandingsTable.StandingsLists[0].DriverStandings {
		fmt.Printf("%-20s %s\n",
			standing.Driver.FamilyName,
			standing.Points,
		)
	}
}
