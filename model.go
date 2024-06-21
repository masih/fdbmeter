package fdbmeter

type Status struct {
	Client struct {
		ClusterFile struct {
			Path     string `json:"path" fdbmeter:"skip"`
			UpToDate bool   `json:"up_to_date"`
		} `json:"cluster_file"`
		Coordinators struct {
			Coordinators []struct {
				Address   string `json:"address"`
				Protocol  string `json:"protocol"`
				Reachable bool   `json:"reachable"`
			} `json:"coordinators"`
			QuorumReachable bool `json:"quorum_reachable"`
		} `json:"coordinators"`
		DatabaseStatus struct {
			Available bool `json:"available"`
			Healthy   bool `json:"healthy"`
		} `json:"database_status"`
		Messages  []any   `json:"messages"`
		Timestamp float64 `json:"timestamp"`
	} `json:"client"`
	Cluster struct {
		ActivePrimaryDc string `json:"active_primary_dc"`
		ActiveTssCount  int    `json:"active_tss_count"`
		BounceImpact    struct {
			CanCleanBounce bool `json:"can_clean_bounce"`
		} `json:"bounce_impact"`
		Clients struct {
			Count             int `json:"count"`
			SupportedVersions []struct {
				ClientVersion    string `json:"client_version"`
				ConnectedClients []struct {
					Address  string `json:"address"`
					LogGroup string `json:"log_group"`
				} `json:"connected_clients"`
				Count              int `json:"count"`
				MaxProtocolClients []struct {
					Address  string `json:"address"`
					LogGroup string `json:"log_group"`
				} `json:"max_protocol_clients"`
				MaxProtocolCount int    `json:"max_protocol_count"`
				ProtocolVersion  string `json:"protocol_version"`
				SourceVersion    string `json:"source_version"`
			} `json:"supported_versions"`
		} `json:"clients"`
		ClusterControllerTimestamp float64 `json:"cluster_controller_timestamp"`
		Configuration              struct {
			BackupWorkerEnabled            int    `json:"backup_worker_enabled"`
			BlobGranulesEnabled            int    `json:"blob_granules_enabled"`
			CommitProxies                  int    `json:"commit_proxies"`
			CoordinatorsCount              int    `json:"coordinators_count"`
			ExcludedServers                []any  `json:"excluded_servers"`
			GrvProxies                     int    `json:"grv_proxies"`
			LogRouters                     int    `json:"log_routers"`
			LogSpill                       int    `json:"log_spill"`
			Logs                           int    `json:"logs"`
			PerpetualStorageWiggle         int    `json:"perpetual_storage_wiggle"`
			PerpetualStorageWiggleLocality string `json:"perpetual_storage_wiggle_locality"`
			Proxies                        int    `json:"proxies"`
			RedundancyMode                 string `json:"redundancy_mode"`
			RemoteLogs                     int    `json:"remote_logs"`
			Resolvers                      int    `json:"resolvers"`
			StorageEngine                  string `json:"storage_engine"`
			StorageMigrationType           string `json:"storage_migration_type"`
			TenantMode                     string `json:"tenant_mode"`
			UsableRegions                  int    `json:"usable_regions"`
		} `json:"configuration"`
		ConnectionString string `json:"connection_string" fdbmeter:"skip"`
		Data             struct {
			AveragePartitionSizeBytes             int   `json:"average_partition_size_bytes"`
			LeastOperatingSpaceBytesLogServer     int64 `json:"least_operating_space_bytes_log_server"`
			LeastOperatingSpaceBytesStorageServer int64 `json:"least_operating_space_bytes_storage_server"`
			MovingData                            struct {
				HighestPriority   int   `json:"highest_priority"`
				InFlightBytes     int64 `json:"in_flight_bytes"`
				InQueueBytes      int64 `json:"in_queue_bytes"`
				TotalWrittenBytes int64 `json:"total_written_bytes"`
			} `json:"moving_data"`
			PartitionsCount int `json:"partitions_count"`
			State           struct {
				Description          string `json:"description"`
				Healthy              bool   `json:"healthy"`
				MinReplicasRemaining int    `json:"min_replicas_remaining"`
				Name                 string `json:"name"`
			} `json:"state"`
			SystemKvSizeBytes int `json:"system_kv_size_bytes"`
			TeamTrackers      []struct {
				InFlightBytes int64 `json:"in_flight_bytes"`
				Primary       bool  `json:"primary"`
				State         struct {
					Description          string `json:"description"`
					Healthy              bool   `json:"healthy"`
					MinReplicasRemaining int    `json:"min_replicas_remaining"`
					Name                 string `json:"name"`
				} `json:"state"`
				UnhealthyServers int `json:"unhealthy_servers"`
			} `json:"team_trackers"`
			TotalDiskUsedBytes int64 `json:"total_disk_used_bytes"`
			TotalKvSizeBytes   int64 `json:"total_kv_size_bytes"`
		} `json:"data"`
		DatabaseAvailable bool `json:"database_available"`
		DatabaseLockState struct {
			Locked bool `json:"locked"`
		} `json:"database_lock_state"`
		DatacenterLag struct {
			Seconds  int `json:"seconds"`
			Versions int `json:"versions"`
		} `json:"datacenter_lag"`
		DegradedProcesses int `json:"degraded_processes"`
		FaultTolerance    struct {
			MaxZoneFailuresWithoutLosingAvailability int `json:"max_zone_failures_without_losing_availability"`
			MaxZoneFailuresWithoutLosingData         int `json:"max_zone_failures_without_losing_data"`
		} `json:"fault_tolerance"`
		FullReplication         bool  `json:"full_replication"`
		Generation              int   `json:"generation"`
		IncompatibleConnections []any `json:"incompatible_connections"`
		LatencyProbe            struct {
			CommitSeconds                            float64 `json:"commit_seconds"`
			ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
			ReadSeconds                              float64 `json:"read_seconds"`
			TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
		} `json:"latency_probe"`
		Layers struct {
			Error string `json:"_error"`
			Valid bool   `json:"_valid"`
		} `json:"layers"`
		Logs []struct {
			BeginVersion      int64 `json:"begin_version"`
			Current           bool  `json:"current"`
			Epoch             int   `json:"epoch"`
			LogFaultTolerance int   `json:"log_fault_tolerance"`
			LogInterfaces     []struct {
				Address string `json:"address"`
				Healthy bool   `json:"healthy"`
				Id      string `json:"id"`
			} `json:"log_interfaces"`
			LogReplicationFactor int  `json:"log_replication_factor"`
			LogWriteAntiQuorum   int  `json:"log_write_anti_quorum"`
			PossiblyLosingData   bool `json:"possibly_losing_data"`
		} `json:"logs"`
		Machines map[string]struct {
			Address             string `json:"address"`
			ContributingWorkers int    `json:"contributing_workers"`
			Cpu                 struct {
				LogicalCoreUtilization float64 `json:"logical_core_utilization"`
			} `json:"cpu"`
			Excluded bool `json:"excluded"`
			Locality struct {
				DnsName    string `json:"dns_name"`
				InstanceId string `json:"instance_id"`
				Machineid  string `json:"machineid"`
				Processid  string `json:"processid"`
				Zoneid     string `json:"zoneid"`
			} `json:"locality"`
			MachineId string `json:"machine_id"`
			Memory    struct {
				CommittedBytes int   `json:"committed_bytes"`
				FreeBytes      int64 `json:"free_bytes"`
				TotalBytes     int64 `json:"total_bytes"`
			} `json:"memory"`
			Network struct {
				MegabitsReceived struct {
					Hz float64 `json:"hz"`
				} `json:"megabits_received"`
				MegabitsSent struct {
					Hz float64 `json:"hz"`
				} `json:"megabits_sent"`
				TcpSegmentsRetransmitted struct {
					Hz float64 `json:"hz"`
				} `json:"tcp_segments_retransmitted"`
			} `json:"network"`
		} `json:"machines" fdbmeter:"key=name"`
		Messages []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"messages"`
		PageCache struct {
			LogHitRate     float64 `json:"log_hit_rate"`
			StorageHitRate float64 `json:"storage_hit_rate"`
		} `json:"page_cache"`
		Processes map[string]struct {
			Address     string `json:"address"`
			ClassSource string `json:"class_source"`
			ClassType   string `json:"class_type"`
			CommandLine string `json:"command_line" fdbmeter:"skip"`
			Cpu         struct {
				UsageCores float64 `json:"usage_cores"`
			} `json:"cpu"`
			Disk struct {
				Busy      float64 `json:"busy"`
				FreeBytes int64   `json:"free_bytes"`
				Reads     struct {
					Counter int     `json:"counter"`
					Hz      float64 `json:"hz"`
					Sectors float64 `json:"sectors"`
				} `json:"reads"`
				TotalBytes int64 `json:"total_bytes"`
				Writes     struct {
					Counter int     `json:"counter"`
					Hz      float64 `json:"hz"`
					Sectors float64 `json:"sectors"`
				} `json:"writes"`
			} `json:"disk"`
			Excluded    bool   `json:"excluded"`
			FaultDomain string `json:"fault_domain"`
			Locality    struct {
				DnsName    string `json:"dns_name"`
				InstanceId string `json:"instance_id"`
				Machineid  string `json:"machineid"`
				ProcessId  string `json:"process_id"`
				Processid  string `json:"processid"`
				Zoneid     string `json:"zoneid"`
			} `json:"locality"`
			MachineId string `json:"machine_id"`
			Memory    struct {
				AvailableBytes        int64 `json:"available_bytes"`
				LimitBytes            int64 `json:"limit_bytes"`
				RssBytes              int64 `json:"rss_bytes"`
				UnusedAllocatedMemory int   `json:"unused_allocated_memory"`
				UsedBytes             int64 `json:"used_bytes"`
			} `json:"memory"`
			Messages []any `json:"messages"`
			Network  struct {
				ConnectionErrors struct {
					Hz float64 `json:"hz"`
				} `json:"connection_errors"`
				ConnectionsClosed struct {
					Hz float64 `json:"hz"`
				} `json:"connections_closed"`
				ConnectionsEstablished struct {
					Hz float64 `json:"hz"`
				} `json:"connections_established"`
				CurrentConnections int `json:"current_connections"`
				MegabitsReceived   struct {
					Hz float64 `json:"hz"`
				} `json:"megabits_received"`
				MegabitsSent struct {
					Hz float64 `json:"hz"`
				} `json:"megabits_sent"`
				TlsPolicyFailures struct {
					Hz float64 `json:"hz"`
				} `json:"tls_policy_failures"`
			} `json:"network"`
			Roles []struct {
				BytesQueried struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"bytes_queried"`
				DataLag struct {
					Seconds  float64 `json:"seconds"`
					Versions int     `json:"versions"`
				} `json:"data_lag"`
				DataVersion   int64 `json:"data_version"`
				DurabilityLag struct {
					Seconds  float64 `json:"seconds"`
					Versions int     `json:"versions"`
				} `json:"durability_lag"`
				DurableBytes struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"durable_bytes"`
				DurableVersion  int64 `json:"durable_version"`
				FetchedVersions struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"fetched_versions"`
				FetchesFromLogs struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"fetches_from_logs"`
				FinishedQueries struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"finished_queries"`
				Id         string `json:"id"`
				InputBytes struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"input_bytes"`
				KeysQueried struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"keys_queried"`
				KvstoreAvailableBytes int64 `json:"kvstore_available_bytes"`
				KvstoreFreeBytes      int64 `json:"kvstore_free_bytes"`
				KvstoreInlineKeys     int   `json:"kvstore_inline_keys"`
				KvstoreTotalBytes     int64 `json:"kvstore_total_bytes"`
				KvstoreTotalNodes     int   `json:"kvstore_total_nodes"`
				KvstoreTotalSize      int   `json:"kvstore_total_size"`
				KvstoreUsedBytes      int64 `json:"kvstore_used_bytes"`
				LocalRate             int   `json:"local_rate"`
				LowPriorityQueries    struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"low_priority_queries"`
				MutationBytes struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"mutation_bytes"`
				Mutations struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"mutations"`
				QueryQueueMax         int `json:"query_queue_max"`
				ReadLatencyStatistics struct {
					Count  int     `json:"count"`
					Max    float64 `json:"max"`
					Mean   float64 `json:"mean"`
					Median float64 `json:"median"`
					Min    float64 `json:"min"`
					P25    float64 `json:"p25"`
					P90    float64 `json:"p90"`
					P95    float64 `json:"p95"`
					P99    float64 `json:"p99"`
					P999   float64 `json:"p99.9"`
				} `json:"read_latency_statistics"`
				Role            string `json:"role"`
				StorageMetadata struct {
					CreatedTimeDatetime  string  `json:"created_time_datetime"`
					CreatedTimeTimestamp float64 `json:"created_time_timestamp"`
				} `json:"storage_metadata"`
				StoredBytes  int64 `json:"stored_bytes"`
				TotalQueries struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"total_queries"`
			} `json:"roles"`
			RunLoopBusy   float64 `json:"run_loop_busy"`
			UptimeSeconds float64 `json:"uptime_seconds"`
			Version       string  `json:"version"`
		} `json:"processes" fdbmeter:"key=process_id"`
		ProtocolVersion string `json:"protocol_version"`
		Qos             struct {
			BatchPerformanceLimitedBy struct {
				Description    string `json:"description"`
				Name           string `json:"name"`
				ReasonId       int    `json:"reason_id"`
				ReasonServerId string `json:"reason_server_id"`
			} `json:"batch_performance_limited_by"`
			BatchReleasedTransactionsPerSecond float64 `json:"batch_released_transactions_per_second"`
			BatchTransactionsPerSecondLimit    float64 `json:"batch_transactions_per_second_limit"`
			LimitingDataLagStorageServer       struct {
				Seconds  float64 `json:"seconds"`
				Versions int     `json:"versions"`
			} `json:"limiting_data_lag_storage_server"`
			LimitingDurabilityLagStorageServer struct {
				Seconds  float64 `json:"seconds"`
				Versions int     `json:"versions"`
			} `json:"limiting_durability_lag_storage_server"`
			LimitingQueueBytesStorageServer int `json:"limiting_queue_bytes_storage_server"`
			PerformanceLimitedBy            struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				ReasonId    int    `json:"reason_id"`
			} `json:"performance_limited_by"`
			ReleasedTransactionsPerSecond float64 `json:"released_transactions_per_second"`
			ThrottledTags                 struct {
				Auto struct {
					BusyRead        int `json:"busy_read"`
					BusyWrite       int `json:"busy_write"`
					Count           int `json:"count"`
					RecommendedOnly int `json:"recommended_only"`
				} `json:"auto"`
				Manual struct {
					Count int `json:"count"`
				} `json:"manual"`
			} `json:"throttled_tags"`
			TransactionsPerSecondLimit float64 `json:"transactions_per_second_limit"`
			WorstDataLagStorageServer  struct {
				Seconds  float64 `json:"seconds"`
				Versions int     `json:"versions"`
			} `json:"worst_data_lag_storage_server"`
			WorstDurabilityLagStorageServer struct {
				Seconds  float64 `json:"seconds"`
				Versions int     `json:"versions"`
			} `json:"worst_durability_lag_storage_server"`
			WorstQueueBytesLogServer     int `json:"worst_queue_bytes_log_server"`
			WorstQueueBytesStorageServer int `json:"worst_queue_bytes_storage_server"`
		} `json:"qos"`
		RecoveryState struct {
			ActiveGenerations         int     `json:"active_generations"`
			Description               string  `json:"description"`
			Name                      string  `json:"name"`
			SecondsSinceLastRecovered float64 `json:"seconds_since_last_recovered"`
		} `json:"recovery_state"`
		Workload struct {
			Bytes struct {
				Read struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read"`
				Written struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"written"`
			} `json:"bytes"`
			Keys struct {
				Read struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read"`
			} `json:"keys"`
			Operations struct {
				LocationRequests struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"location_requests"`
				LowPriorityReads struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"low_priority_reads"`
				MemoryErrors struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"memory_errors"`
				ReadRequests struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"read_requests"`
				Reads struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"reads"`
				Writes struct {
					Counter   int64   `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"writes"`
			} `json:"operations"`
			Transactions struct {
				Committed struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"committed"`
				Conflicted struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"conflicted"`
				RejectedForQueuedTooLong struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"rejected_for_queued_too_long"`
				Started struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started"`
				StartedBatchPriority struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_batch_priority"`
				StartedDefaultPriority struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_default_priority"`
				StartedImmediatePriority struct {
					Counter   int     `json:"counter"`
					Hz        float64 `json:"hz"`
					Roughness float64 `json:"roughness"`
				} `json:"started_immediate_priority"`
			} `json:"transactions"`
		} `json:"workload"`
	} `json:"cluster"`
}
