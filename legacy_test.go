package sc2

import (
	"encoding/json"
	"testing"
)

func TestGetLegacyProfile(t *testing.T) {
	filepath := "test_fixtures/legacy_profile_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	prof, err := client.GetLegacyProfile(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	prof_b, err := json.Marshal(prof)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(prof_b, fixture) {
		writeTestDebugOutput(prof_b, fixture, "test_get_legacy_profile")
		t.Error("Fields for target struct LegacyProfile does not match the provided fixture")
	}
}

func TestGetLegacyLadders(t *testing.T) {
	filepath := "test_fixtures/legacy_ladders_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	ladd, err := client.GetLegacyLadders(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	ladd_b, err := json.Marshal(ladd)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(ladd_b, fixture) {
		writeTestDebugOutput(ladd_b, fixture, "test_get_legacy_ladders")
		t.Error("Fields for target struct LegacyLadders does not match the provided fixture")
	}
}

func TestGetLegacyMatches(t *testing.T) {
	filepath := "test_fixtures/legacy_matches_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	match, err := client.GetLegacyMatches(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	match_b, err := json.Marshal(match)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(match_b, fixture) {
		writeTestDebugOutput(match_b, fixture, "test_get_legacy_matches")
		t.Error("Fields for target struct LegacyMatches does not match the provided fixture")
	}
}

func TestGetLegacyAchievements(t *testing.T) {
	filepath := "test_fixtures/legacy_achievements_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	achi, err := client.GetLegacyAchievements(1)
	if err != nil {
		t.Fatal(err)
	}
	achi_b, err := json.Marshal(achi)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(achi_b, fixture) {
		writeTestDebugOutput(achi_b, fixture, "test_get_legacy_achievements")
		t.Error("Fields for target struct LegacyAchievements does not match the provided fixture")
	}
}

func TestGetLegacyLadder(t *testing.T) {
	filepath := "test_fixtures/legacy_ladder_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	ladd, err := client.GetLegacyLadder(1, 112233)
	if err != nil {
		t.Fatal(err)
	}
	ladd_b, err := json.Marshal(ladd)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(ladd_b, fixture) {
		writeTestDebugOutput(ladd_b, fixture, "test_get_legacy_ladder")
		t.Error("Fields for target struct LegacyLadder does not match the provided fixture")
	}
}

func GetLegacyRewards(t *testing.T) {
	filepath := "test_fixtures/legacy_rewards_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	rewa, err := client.GetLegacyRewards(1)
	if err != nil {
		t.Fatal(err)
	}
	rewa_b, err := json.Marshal(rewa)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(rewa_b, fixture) {
		writeTestDebugOutput(rewa_b, fixture, "test_get_legacy_rewards")
		t.Error("Fields for target struct LegacyRewards does not match the provided fixture")
	}
}
