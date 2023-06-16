package fdbmeter

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStatusDecode(t *testing.T) {
	tests := []struct {
		jsonFile string
	}{
		{jsonFile: "testdata/status1.json"},
		{jsonFile: "testdata/status2.json"},
	}
	for _, test := range tests {
		t.Run(test.jsonFile, func(t *testing.T) {
			f, err := os.Open(test.jsonFile)
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					t.Fatal(err)
				}
			}()
			var status Status
			if err := json.NewDecoder(f).Decode(&status); err != nil {
				t.Fatal(err)
			}
		})
	}
}
