package fdbmeter

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

func TestMetricsPopulatesValues(t *testing.T) {
	tests := []struct {
		jsonFile string
		want     map[string][]observable
	}{
		{
			jsonFile: "testdata/status1.json",
			want: map[string][]observable{
				"client_cluster_file_up_to_date": {
					{
						int64Value: 1,
						attrs:      attribute.NewSet(),
					},
				},
				"cluster_data_state_healthy": {
					{int64Value: 0, attrs: attribute.NewSet(attribute.String("description", "Only one replica remains of some data"), attribute.String("name", "healing"))},
				},
			},
		},
		{
			jsonFile: "testdata/status2.json",
			want: map[string][]observable{
				"cluster_workload_operations_writes_hz": {
					{
						float64Value: 342887,
						attrs:        attribute.NewSet(),
					},
				},
				"cluster_latency_probe_immediate_priority_transaction_start_seconds": {
					{
						float64Value: 0.6852229999999999,
						attrs:        attribute.NewSet(),
					},
				},
			},
		},
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

			subject, err := NewMetrics()
			if err != nil {
				t.Fatal(err)
			}
			subject.notifyStatus(context.Background(), status)

			for k, want := range test.want {

				got := subject.observables[k]
				if len(got) != len(want) {
					t.Fatal()
				}

				for i, target := range got {
					if !reflect.DeepEqual(want[i], target) {
						t.Fatal(want[i], " got ", target)
					}
				}

			}
		})
	}
}
