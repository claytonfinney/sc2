package sc2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	url := fmt.Sprintf("https://%s.battle.net/oauth/token?grant_type=client_credentials", region)
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
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(i, &a)
	if err != nil {
		log.Println(string(i))
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
	i, err := ioutil.ReadAll(resp.Body)
	log.Println(string(i))
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(i, &payload)
	if err != nil {
		log.Println(string(i))
		log.Fatal(err)
	}
	if err != nil {
		return err
	}
	return nil
}
