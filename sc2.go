package sc2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct to Unmarshal auth token payload
type AuthToken struct {
	Token string `json:"access_token"`
	Error string `json:"error"`
}

// Main point of interaction with the SC2 web API
type Client struct {
	BaseUrl string
	Client  string
	Secret  string
	Token   string
}

func NewClient(region string, client string, secret string) *Client {
	url := fmt.Sprintf("https://%s.battle.net/ouath/token&grant_type=client_credentials", region)
	h := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(client, secret)
	resp, err := h.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	a := AuthToken{}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&a)
	if err != nil {
		log.Fatal(err)
	}

	if a.Error != "" {
		log.Fatalf("Request to %s resulted in status code %s, please check your credentials", url, resp.StatusCode)
	}

	base := fmt.Sprintf("https://%s.api.blizzard.com", region)
	return &Client{BaseUrl: base, Client: client, Secret: secret, Token: a.Token}
}

func (c *Client) Get(uri string, payload interface{}) error {
	full := fmt.Sprintf("%s?access_token=%s", uri, c.Token)
	resp, err := http.Get(full)
	if err != nil {
		return err
	}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(payload)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	c := "YOUR-CLIENT-ID"
	s := "YOUR-SECRET-ID"
	client := NewClient("us", c, s)

	szn, err := client.GetSeason(1)
	if err != nil {
		log.Fatal(err)
	}
	curr := szn.SeasonId // corresponding Blizzard JSON field is seasonId
	// seasonId, queueId = 201 for LotV, teamType = 0 for 1v1, leagueId = 6 for Grandmaster)
	d, err := client.GetLeagueData(curr, 201, 0, 6)
	if err != nil {
		log.Fatal(err)
	}
	id := d.Tier[0].Division.Id
	grandmasters, err := client.GetLegacyLadder(1, id) // Currently the way to fetch ladder data without a unique player ID, may change soon
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grandmasters.Team[0].Member[0].LegacyLink.Name) // Gets the first place Grandmaster!
}
