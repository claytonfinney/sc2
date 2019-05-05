package sc2

import (
	"encoding/json"
	"testing"
)

func TestGetLadderSummary(t *testing.T) {
	filepath := "test_fixtures/ladder_summary_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	summ, err := client.GetLadderSummary(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	summ_b, err := json.Marshal(summ)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(summ_b, fixture) {
		writeTestDebugOutput(summ_b, fixture, "test_get_ladder_summary")
		t.Error("Fields for target struct LadderSummary does not match the provided fixture")
	}
}

func TestGetLadder(t *testing.T) {
	filepath := "test_fixtures/ladder_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	ladd, err := client.GetLadder(1, 1, 123456, 000001)
	if err != nil {
		t.Fatal(err)
	}
	ladd_b, err := json.Marshal(ladd)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(ladd_b, fixture) {
		writeTestDebugOutput(ladd_b, fixture, "test_get_ladder")
		t.Error("Fields for target struct Ladder does not match the provided fixture")
	}
}

// Bit of a rehash of TestGetLadder, maybe do a test to support the length being > 100 or something?
func TestGetGrandmaster(t *testing.T) {
	filepath := "test_fixtures/ladder_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	ladd, err := client.GetGrandmaster(1)
	if err != nil {
		t.Fatal(err)
	}
	ladd_b, err := json.Marshal(ladd)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(ladd_b, fixture) {
		writeTestDebugOutput(ladd_b, fixture, "test_get_grandmaster")
		t.Error("Fields for target struct Ladder does not match the provided fixture")
	}
}

func TestGetSeason(t *testing.T) {
	filepath := "test_fixtures/season_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	seas, err := client.GetSeason(1)
	if err != nil {
		t.Fatal(err)
	}
	seas_b, err := json.Marshal(seas)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(seas_b, fixture) {
		writeTestDebugOutput(seas_b, fixture, "test_get_season")
		t.Error("Fields for target struct Season does not match the provided fixture")
	}
	if seas.SeasonId != 39 {
		t.Error("Season.SeasonId field did not match desired 37")
	}
}
