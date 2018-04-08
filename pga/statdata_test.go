package pga

import (
	"testing"
)

func TestGetCurrentTournamentt(t *testing.T) {

	m := GetCurrentTournament()

	headerI := m["header"]
	header := headerI.(map[string]interface{})

	if header["version"] != "1.0" {
		t.Error("Unsupported version at endpoint")
	}

	if m["tc"] != "r" {
		t.Fail()
	}
}
