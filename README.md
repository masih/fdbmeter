# :detective: `fdbmeter`

`fdbmeter` is a FoundationDB Prometheus metrics exporter that:

* Works with the latest FoundationDB release, currently `7.1.33`
* Uses reflection to automatically generate and export Prometheus metrics
    * This means changes to FoundationDB status schema only need Golang struct changes.
* Automatically tags the metrics with relevant attributes extracted from textual status json fields.
* Periodically fetches and caches FoundationDB status.
* Exposes a REST endpoint at `/status` to serve the raw JSON for further debugging.
* Supports custom struct tags to customize metric reporting using the key `fdbmeter`; see [`model`](model.go).

## Metrics

Below is the current list of metrics exposed by `fdbmeter`:

```text
fdb_client_cluster_file_up_to_date
fdb_client_coordinators_coordinators_reachable
fdb_client_coordinators_quorum_reachable
fdb_client_database_status_available
fdb_client_database_status_healthy
fdb_client_timestamp
fdb_cluster_active_tss_count
fdb_cluster_bounce_impact_can_clean_bounce
fdb_cluster_clients_count
fdb_cluster_clients_supported_versions_count
fdb_cluster_clients_supported_versions_max_protocol_count
fdb_cluster_cluster_controller_timestamp
fdb_cluster_configuration_backup_worker_enabled
fdb_cluster_configuration_blob_granules_enabled
fdb_cluster_configuration_commit_proxies
fdb_cluster_configuration_coordinators_count
fdb_cluster_configuration_grv_proxies
fdb_cluster_configuration_log_routers
fdb_cluster_configuration_log_spill
fdb_cluster_configuration_logs
fdb_cluster_configuration_perpetual_storage_wiggle
fdb_cluster_configuration_proxies
fdb_cluster_configuration_remote_logs
fdb_cluster_configuration_resolvers
fdb_cluster_configuration_usable_regions
fdb_cluster_data_average_partition_size_bytes
fdb_cluster_data_least_operating_space_bytes_log_server
fdb_cluster_data_least_operating_space_bytes_storage_server
fdb_cluster_data_moving_data_highest_priority
fdb_cluster_data_moving_data_in_flight_bytes
fdb_cluster_data_moving_data_in_queue_bytes
fdb_cluster_data_moving_data_total_written_bytes
fdb_cluster_data_partitions_count
fdb_cluster_data_state_healthy
fdb_cluster_data_state_min_replicas_remaining
fdb_cluster_data_system_kv_size_bytes
fdb_cluster_data_team_trackers_in_flight_bytes
fdb_cluster_data_team_trackers_primary
fdb_cluster_data_team_trackers_state_healthy
fdb_cluster_data_team_trackers_state_min_replicas_remaining
fdb_cluster_data_team_trackers_unhealthy_servers
fdb_cluster_data_total_disk_used_bytes
fdb_cluster_data_total_kv_size_bytes
fdb_cluster_database_available
fdb_cluster_database_lock_state_locked
fdb_cluster_datacenter_lag_seconds
fdb_cluster_datacenter_lag_versions
fdb_cluster_degraded_processes
fdb_cluster_fault_tolerance_max_zone_failures_without_losing_availability
fdb_cluster_fault_tolerance_max_zone_failures_without_losing_data
fdb_cluster_full_replication
fdb_cluster_generation
fdb_cluster_latency_probe_commit_seconds
fdb_cluster_latency_probe_immediate_priority_transaction_start_seconds
fdb_cluster_latency_probe_read_seconds
fdb_cluster_latency_probe_transaction_start_seconds
fdb_cluster_layers__valid
fdb_cluster_logs_begin_version
fdb_cluster_logs_current
fdb_cluster_logs_epoch
fdb_cluster_logs_log_fault_tolerance
fdb_cluster_logs_log_interfaces_healthy
fdb_cluster_logs_log_replication_factor
fdb_cluster_logs_log_write_anti_quorum
fdb_cluster_logs_possibly_losing_data
fdb_cluster_machines_contributing_workers
fdb_cluster_machines_cpu_logical_core_utilization
fdb_cluster_machines_excluded
fdb_cluster_machines_memory_committed_bytes
fdb_cluster_machines_memory_free_bytes
fdb_cluster_machines_memory_total_bytes
fdb_cluster_machines_network_megabits_received_hz
fdb_cluster_machines_network_megabits_sent_hz
fdb_cluster_machines_network_tcp_segments_retransmitted_hz
fdb_cluster_page_cache_log_hit_rate
fdb_cluster_page_cache_storage_hit_rate
fdb_cluster_processes_cpu_usage_cores
fdb_cluster_processes_disk_busy
fdb_cluster_processes_disk_free_bytes
fdb_cluster_processes_disk_reads_counter
fdb_cluster_processes_disk_reads_hz
fdb_cluster_processes_disk_reads_sectors
fdb_cluster_processes_disk_total_bytes
fdb_cluster_processes_disk_writes_counter
fdb_cluster_processes_disk_writes_hz
fdb_cluster_processes_disk_writes_sectors
fdb_cluster_processes_excluded
fdb_cluster_processes_memory_available_bytes
fdb_cluster_processes_memory_limit_bytes
fdb_cluster_processes_memory_rss_bytes
fdb_cluster_processes_memory_unused_allocated_memory
fdb_cluster_processes_memory_used_bytes
fdb_cluster_processes_network_connection_errors_hz
fdb_cluster_processes_network_connections_closed_hz
fdb_cluster_processes_network_connections_established_hz
fdb_cluster_processes_network_current_connections
fdb_cluster_processes_network_megabits_received_hz
fdb_cluster_processes_network_megabits_sent_hz
fdb_cluster_processes_network_tls_policy_failures_hz
fdb_cluster_processes_roles_bytes_queried_counter
fdb_cluster_processes_roles_bytes_queried_hz
fdb_cluster_processes_roles_bytes_queried_roughness
fdb_cluster_processes_roles_data_lag_seconds
fdb_cluster_processes_roles_data_lag_versions
fdb_cluster_processes_roles_data_version
fdb_cluster_processes_roles_durability_lag_seconds
fdb_cluster_processes_roles_durability_lag_versions
fdb_cluster_processes_roles_durable_bytes_counter
fdb_cluster_processes_roles_durable_bytes_hz
fdb_cluster_processes_roles_durable_bytes_roughness
fdb_cluster_processes_roles_durable_version
fdb_cluster_processes_roles_fetched_versions_counter
fdb_cluster_processes_roles_fetched_versions_hz
fdb_cluster_processes_roles_fetched_versions_roughness
fdb_cluster_processes_roles_fetches_from_logs_counter
fdb_cluster_processes_roles_fetches_from_logs_hz
fdb_cluster_processes_roles_fetches_from_logs_roughness
fdb_cluster_processes_roles_finished_queries_counter
fdb_cluster_processes_roles_finished_queries_hz
fdb_cluster_processes_roles_finished_queries_roughness
fdb_cluster_processes_roles_input_bytes_counter
fdb_cluster_processes_roles_input_bytes_hz
fdb_cluster_processes_roles_input_bytes_roughness
fdb_cluster_processes_roles_keys_queried_counter
fdb_cluster_processes_roles_keys_queried_hz
fdb_cluster_processes_roles_keys_queried_roughness
fdb_cluster_processes_roles_kvstore_available_bytes
fdb_cluster_processes_roles_kvstore_free_bytes
fdb_cluster_processes_roles_kvstore_inline_keys
fdb_cluster_processes_roles_kvstore_total_bytes
fdb_cluster_processes_roles_kvstore_total_nodes
fdb_cluster_processes_roles_kvstore_total_size
fdb_cluster_processes_roles_kvstore_used_bytes
fdb_cluster_processes_roles_local_rate
fdb_cluster_processes_roles_low_priority_queries_counter
fdb_cluster_processes_roles_low_priority_queries_hz
fdb_cluster_processes_roles_low_priority_queries_roughness
fdb_cluster_processes_roles_mutation_bytes_counter
fdb_cluster_processes_roles_mutation_bytes_hz
fdb_cluster_processes_roles_mutation_bytes_roughness
fdb_cluster_processes_roles_mutations_counter
fdb_cluster_processes_roles_mutations_hz
fdb_cluster_processes_roles_mutations_roughness
fdb_cluster_processes_roles_query_queue_max
fdb_cluster_processes_roles_read_latency_statistics_count
fdb_cluster_processes_roles_read_latency_statistics_max
fdb_cluster_processes_roles_read_latency_statistics_mean
fdb_cluster_processes_roles_read_latency_statistics_median
fdb_cluster_processes_roles_read_latency_statistics_min
fdb_cluster_processes_roles_read_latency_statistics_p25
fdb_cluster_processes_roles_read_latency_statistics_p90
fdb_cluster_processes_roles_read_latency_statistics_p95
fdb_cluster_processes_roles_read_latency_statistics_p99
fdb_cluster_processes_roles_read_latency_statistics_p99.9
fdb_cluster_processes_roles_storage_metadata_created_time_timestamp
fdb_cluster_processes_roles_stored_bytes
fdb_cluster_processes_roles_total_queries_counter
fdb_cluster_processes_roles_total_queries_hz
fdb_cluster_processes_roles_total_queries_roughness
fdb_cluster_processes_run_loop_busy
fdb_cluster_processes_uptime_seconds
fdb_cluster_qos_batch_performance_limited_by_reason_id
fdb_cluster_qos_batch_released_transactions_per_second
fdb_cluster_qos_batch_transactions_per_second_limit
fdb_cluster_qos_limiting_data_lag_storage_server_seconds
fdb_cluster_qos_limiting_data_lag_storage_server_versions
fdb_cluster_qos_limiting_durability_lag_storage_server_seconds
fdb_cluster_qos_limiting_durability_lag_storage_server_versions
fdb_cluster_qos_limiting_queue_bytes_storage_server
fdb_cluster_qos_performance_limited_by_reason_id
fdb_cluster_qos_released_transactions_per_second
fdb_cluster_qos_throttled_tags_auto_busy_read
fdb_cluster_qos_throttled_tags_auto_busy_write
fdb_cluster_qos_throttled_tags_auto_count
fdb_cluster_qos_throttled_tags_auto_recommended_only
fdb_cluster_qos_throttled_tags_manual_count
fdb_cluster_qos_transactions_per_second_limit
fdb_cluster_qos_worst_data_lag_storage_server_seconds
fdb_cluster_qos_worst_data_lag_storage_server_versions
fdb_cluster_qos_worst_durability_lag_storage_server_seconds
fdb_cluster_qos_worst_durability_lag_storage_server_versions
fdb_cluster_qos_worst_queue_bytes_log_server
fdb_cluster_qos_worst_queue_bytes_storage_server
fdb_cluster_recovery_state_active_generations
fdb_cluster_recovery_state_seconds_since_last_recovered
fdb_cluster_workload_bytes_read_counter
fdb_cluster_workload_bytes_read_hz
fdb_cluster_workload_bytes_read_roughness
fdb_cluster_workload_bytes_written_counter
fdb_cluster_workload_bytes_written_hz
fdb_cluster_workload_bytes_written_roughness
fdb_cluster_workload_keys_read_counter
fdb_cluster_workload_keys_read_hz
fdb_cluster_workload_keys_read_roughness
fdb_cluster_workload_operations_location_requests_counter
fdb_cluster_workload_operations_location_requests_hz
fdb_cluster_workload_operations_location_requests_roughness
fdb_cluster_workload_operations_low_priority_reads_counter
fdb_cluster_workload_operations_low_priority_reads_hz
fdb_cluster_workload_operations_low_priority_reads_roughness
fdb_cluster_workload_operations_memory_errors_counter
fdb_cluster_workload_operations_memory_errors_hz
fdb_cluster_workload_operations_memory_errors_roughness
fdb_cluster_workload_operations_read_requests_counter
fdb_cluster_workload_operations_read_requests_hz
fdb_cluster_workload_operations_read_requests_roughness
fdb_cluster_workload_operations_reads_counter
fdb_cluster_workload_operations_reads_hz
fdb_cluster_workload_operations_reads_roughness
fdb_cluster_workload_operations_writes_counter
fdb_cluster_workload_operations_writes_hz
fdb_cluster_workload_operations_writes_roughness
fdb_cluster_workload_transactions_committed_counter
fdb_cluster_workload_transactions_committed_hz
fdb_cluster_workload_transactions_committed_roughness
fdb_cluster_workload_transactions_conflicted_counter
fdb_cluster_workload_transactions_conflicted_hz
fdb_cluster_workload_transactions_conflicted_roughness
fdb_cluster_workload_transactions_rejected_for_queued_too_long_counter
fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz
fdb_cluster_workload_transactions_rejected_for_queued_too_long_roughness
fdb_cluster_workload_transactions_started_counter
fdb_cluster_workload_transactions_started_hz
fdb_cluster_workload_transactions_started_roughness
fdb_cluster_workload_transactions_started_batch_priority_counter
fdb_cluster_workload_transactions_started_batch_priority_hz
fdb_cluster_workload_transactions_started_batch_priority_roughness
fdb_cluster_workload_transactions_started_default_priority_counter
fdb_cluster_workload_transactions_started_default_priority_hz
fdb_cluster_workload_transactions_started_default_priority_roughness
fdb_cluster_workload_transactions_started_immediate_priority_counter
fdb_cluster_workload_transactions_started_immediate_priority_hz
fdb_cluster_workload_transactions_started_immediate_priority_roughness
```

Additionally, the following metrics expose information about the metric collection health:
```text
fdb_meter_get_status_failure_counter
fdb_meter_get_status_latency_histogram
```

## Prerequisites

The prerequisites below only apply to running `fdbmeter` binary directly on your machine.
Alternatively, run `fdbmeter` via a container runtime environment which only requires pulling the image and a container
runtime of your choice.

* Golang.
* FoundationDB C bindings compatible with API version `710` and above. See
  FoundationDB [release](https://github.com/apple/foundationdb/releases/latest) page.

## Install

To install `fdbmeter` binary directly via Golang, run:

```shell
$ go install github.com/ipni/fdbmeter/cmd/fdbmeter@latest
```

## Usage

```shell
$ fdbmeter --help
Usage of fdbmeter
  -fdbApiVersion int
        The FoundationDB API version. (default 710)
  -fdbClusterFile string
        Path to the FoundationDB cluster file.
  -httpListenAddr string
        The bind address of fdbmeter HTTP server. (default "0.0.0.0:40080")
  -statusRefreshInterval duration
        The interval at which to refresh the FoundationDB status. (default 10s)
```

## License

[SPDX-License-Identifier: Apache-2.0 OR MIT](LICENSE.md)