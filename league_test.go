package sc2

import (
	"encoding/json"
	"testing"
)

func TestGetLeagueData(t *testing.T) {
	filepath := "test_fixtures/league_data_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	data, err := client.GetLeagueData(37, 201, 0, 5)
	if err != nil {
		t.Fatal(err)
	}
	data_b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(data_b, fixture) {
		writeTestDebugOutput(data_b, fixture, "test_get_league_data")
		t.Error("Fields for target struct LeagueData does not match the provided fixture")
	}
}
