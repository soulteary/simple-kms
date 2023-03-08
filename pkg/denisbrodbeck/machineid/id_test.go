package machineid

import "testing"

func TestID(t *testing.T) {
	got, err := ID()
	if err != nil {
		t.Error(err)
	}
	if got == "" {
		t.Error("Got empty machine id")
	}
}
