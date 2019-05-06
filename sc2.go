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

type ClientConfig struct {
	// "us" || "eu" || "kr" || "tw" || "cn"
	Region string
	// A Blizzard API Client ID
	Client string
	// The Blizzard API Secret ID
	Secret string
	// This optional flag allows for us to enable unit testing
	Test bool
	// The URL we're using for our test server
	TestUrl string
}

func getAuthToken(c *ClientConfig) string {
	url := fmt.Sprintf("https://%s.battle.net/oauth/token?grant_type=client_credentials", c.Region)
	h := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(c.Client, c.Secret)
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
		log.Fatal(err)
	}

	if a.Error != "" {
		log.Fatalf("Request to %s resulted in status code %v, please check your credentials", url, resp.StatusCode)
	}

	return a.Token
}

func NewClient(c *ClientConfig) *Client {
	// If we've initialized this client inside a test, we set the client's base URL to be our mock API url
	if c.Test == true {
		var base string
		if c.TestUrl == "" {
			base = "http://localhost:8080/"
		} else {
			base = c.TestUrl
		}
		return &Client{BaseUrl: base, Token: "test_token"}
	}

	// Else, we need to get an auth token from the API and buld a full client
	token := getAuthToken(c)
	base := fmt.Sprintf("https://%s.api.blizzard.com", c.Region)
	return &Client{BaseUrl: base, Client: c.Client, Secret: c.Secret, Token: token}

}

func (c *Client) Get(uri string, payload interface{}) error {
	full := fmt.Sprintf("%s?access_token=%s", uri, c.Token)
	resp, err := http.Get(full)
	if err != nil {
		return err
	}
	i, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(i, &payload)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return err
	}
	return nil
}

func main() {
	c := &ClientConfig{Region: "us", Test: true}
	fmt.Println(c)
}
