package fdbmeter

import (
	"context"
	"encoding/json"
	"fmt"
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
				"cluster_processes_roles_bytes_queried_counter": {
					{
						int64Value: 75292021001,
						attrs: attribute.NewSet(
							attribute.String("active_primary_dc", "abc"),
							attribute.String("address", "20.10.7.15:4509"),
							attribute.String("class_source", "command_line"),
							attribute.String("class_type", "storage"),
							attribute.String("fault_domain", "ip-1-2-7-106"),
							attribute.String("id", "1d43ff854cd58260"),
							attribute.String("machine_id", "ip-1-2-7-106"),
							attribute.String("process_id", "1e5c1cb8cafedd3b4aa335c91c008e5d"),
							attribute.String("protocol_version", "fdb00b071010000"),
							attribute.String("role", "storage"),
							attribute.String("version", "7.1.33"),
						),
					},
				},
				"cluster_data_state_healthy": {
					{int64Value: 0, attrs: attribute.NewSet(
						attribute.String("active_primary_dc", "abc"),
						attribute.String("description", "Only one replica remains of some data"),
						attribute.String("name", "healing"),
						attribute.String("protocol_version", "fdb00b071010000"),
					)},
				},
				"cluster_qos_performance_limited_by_reason_id": {
					{int64Value: 1, attrs: attribute.NewSet(
						attribute.String("active_primary_dc", "abc"),
						attribute.String("description", "Storage server performance (storage queue)."),
						attribute.String("name", "storage_server_write_queue_size"),
						attribute.String("protocol_version", "fdb00b071010000"),
					)},
				},
			},
		},
		{
			jsonFile: "testdata/status2.json",
			want: map[string][]observable{
				"cluster_workload_operations_writes_hz": {
					{
						float64Value: 342887,
						attrs: attribute.NewSet(
							attribute.String("active_primary_dc", ""),
							attribute.String("protocol_version", "fdb00b071010000"),
						),
					},
				},
				"cluster_latency_probe_immediate_priority_transaction_start_seconds": {
					{
						float64Value: 0.6852229999999999,
						attrs: attribute.NewSet(
							attribute.String("active_primary_dc", ""),
							attribute.String("protocol_version", "fdb00b071010000"),
						),
					},
				},
				"cluster_machines_memory_free_bytes": {
					{
						int64Value: 63909851136,
						attrs: attribute.NewSet(
							attribute.String("active_primary_dc", ""),
							attribute.String("address", "20.10.5.192"),
							attribute.String("machine_id", "ip-1-2-5-156"),
							attribute.String("name", "ip-1-2-5-156"),
							attribute.String("protocol_version", "fdb00b071010000"),
						),
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

			// Assert that there are no duplicate attributes for the same metric name.
			for metric, observables := range subject.observables {
				seen := make(map[string]struct{})
				for _, o := range observables {
					key := fmt.Sprintf("%v", o)
					if _, exists := seen[key]; exists {
						t.Fatal("duplicate observable for metric ", metric)
					} else {
						seen[key] = struct{}{}
					}
				}
			}

		WantLoop:
			for k, want := range test.want {

				got := subject.observables[k]
				switch len(want) {
				case 0:
					if len(got) != 0 {
						t.Fatal("expected no observables")
					}
				default:
					for _, target := range got {
						for _, wantObservable := range want {
							if reflect.DeepEqual(wantObservable, target) {
								continue WantLoop
							}
						}
					}
					t.Fatal("did not find expected observable for ", k)
				}
			}
		})
	}
}
