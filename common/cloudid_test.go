package common

import (
	"testing"
)

func TestCloudID(t *testing.T) {
	id := CloudID()
	if id != "AWS" && id != "GCP" && id != "Unknown" {
		t.Errorf("unexpected cloud ID: %s", id)
	}
}
