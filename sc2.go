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
	req, err := http.NewRequest("POST", url)
	if err != nil {
		fmt.Fatal(err)
	}
	req.SetBasicAuth(client, secret)
	resp, err := h.Do(req)
	if err != nil {
		fmt.Fatal(err)
	}
	a := AuthToken{}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(&a)
	if err != nil {
		fmt.Fatal(err)
	}

	if a.Error != "" {
		fmt.Fatalf("Request to %s resulted in status code %s, please check your credentials", url, resp.StatusCode)
	}

	base := fmt.Sprintf("https://%s.api.blizzard.com", region)
	return &Client{BaseUrl: base, Client: client, Secret: secret, Token: a.Token}
}

func (c *Client) Get(uri string, payload *interface{}) {
	full := fmt.Sprintf("%s?access_token=%s", uri, c.AccessToken)
	resp, err := http.Get(env.Config.BnetBaseURL + uri)
	if err != nil {
		return nil, err
	}
	j := json.NewDecoder(resp.Body)
	err = j.Decode(payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}
