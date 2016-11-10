package checks

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/abrander/gansoi/plugins"
)

type (
	mockAgent struct {
		ReturnError bool `json:"return_error"`
	}
)

func (m *mockAgent) Check(result plugins.AgentResult) error {
	if m.ReturnError {
		return errors.New("error")
	}

	result.AddValue("ran", true)

	return nil
}

func init() {
	plugins.RegisterAgent("mock", mockAgent{})
}

func TestCheckJsonInvalid(t *testing.T) {
	cases := []string{
		`{"id": 12}`,
		`{"agent": "nonexisting"}`,
		`{"agent": "mock", "arguments": "wrongtype"}`,
	}

	var check Check
	for _, input := range cases {
		err := json.Unmarshal([]byte(input), &check)

		if err == nil {
			t.Fatalf("Unmarshal did not catch broken json '%s'", input)
		}
	}
}

func TestCheckJson(t *testing.T) {
	input := []byte(`{
    	"id": "tester",
    	"agent": "mock",
    	"arguments": {
    	}
    }`)

	var check Check
	err := json.Unmarshal(input, &check)
	if err != nil {
		t.Fatalf("Unmarshal failed: %s", err.Error())
	}

	if check.ID != "tester" {
		t.Fatalf("ID is not 'test', (got %s)", check.ID)
	}
}

func TestRunCheck(t *testing.T) {
	input := []byte(`{
    	"id": "tester",
    	"agent": "mock",
    	"arguments": {},
        "expressions": ["ran == true"]
    }`)

	var check Check
	err := json.Unmarshal(input, &check)
	if err != nil {
		t.Fatalf("Unmarshal failed: %s", err.Error())
	}

	result := RunCheck(&check)
	if result.Results["ran"] != true {
		t.Fatalf("Check failed to run")
	}
}

func TestRunCheckError(t *testing.T) {
	cases := []string{
		`{"agent": "mock", "arguments": {"return_error": true}, "expressions": ["ran == true"]}`,
		`{"agent": "mock", "arguments": {}, "expressions": ["<<"]}`,
		`{"agent": "mock", "arguments": {}, "expressions": ["nonexisting < 10"]}`,
		`{"agent": "mock", "arguments": {}, "expressions": ["ran == false"]}`,
	}

	var check Check
	for _, input := range cases {
		err := json.Unmarshal([]byte(input), &check)
		if err != nil {
			t.Fatalf("Unmarshal failed: %s", err.Error())
		}

		result := RunCheck(&check)
		if result.Error == "" {
			t.Fatalf("Failed to return error for '%s'", input)
		}
	}
}