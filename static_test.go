package sc2

import (
	"encoding/json"
	"testing"
)

func TestGetStatic(t *testing.T) {
	filepath := "test_fixtures/static_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	stat, err := client.GetStatic(1)
	if err != nil {
		t.Fatal(err)
	}
	stat_b, err := json.Marshal(stat)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(stat_b, fixture) {
		writeTestDebugOutput(stat_b, fixture, "test_get_static")
		t.Error("Fields for target struct Static does not match the provided fixture")
	}
}

func TestGetMetadata(t *testing.T) {
	filepath := "test_fixtures/metadata_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	meta, err := client.GetMetadata(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	meta_b, err := json.Marshal(meta)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(meta_b, fixture) {
		writeTestDebugOutput(meta_b, fixture, "test_get_metadata")
		t.Error("Fields for target struct Metadata does not match the provided fixture")
	}
}

func TestGetProfile(t *testing.T) {
	filepath := "test_fixtures/profile_test_payload.json"
	fixture := getBodyFromFile(filepath)
	client, server := getTestClient(200, fixture)
	defer server.Close()

	// Ensure that all fields made it into the struct by checking equality
	prof, err := client.GetProfile(1, 1, 123456)
	if err != nil {
		t.Fatal(err)
	}
	prof_b, err := json.Marshal(prof)
	if err != nil {
		t.Fatal(err)
	}
	if !byteSlicesMatch(prof_b, fixture) {
		writeTestDebugOutput(prof_b, fixture, "test_get_profile")
		t.Error("Fields for target struct Profile does not match the provided fixture")
	}
}
