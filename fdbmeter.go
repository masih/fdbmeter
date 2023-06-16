package fdbmeter

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const fdbStatusJsonKey = "\xff\xff/status/json"

type FDBMeter struct {
	*options
	db      fdb.Database
	status  atomic.Pointer[Status]
	metrics *Metrics
	server  *http.Server
}

func New(o ...Option) (*FDBMeter, error) {
	opts, err := newOptions(o...)
	if err != nil {
		return nil, err
	}
	metrics, err := NewMetrics()
	if err != nil {
		return nil, err
	}
	return &FDBMeter{
		options: opts,
		metrics: metrics,
	}, nil
}

func (m *FDBMeter) ServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			currentStatus := m.status.Load()
			if currentStatus == nil {
				http.Error(w, "No status found", http.StatusNotFound)
				return
			}
			if err := json.NewEncoder(w).Encode(currentStatus); err != nil {
				log.Printf("Failed to encode status: %v", err)
				http.Error(w, "Failed to encode status as JSON", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}

func (m *FDBMeter) Start(ctx context.Context) error {
	var err error
	if err = fdb.APIVersion(m.fdbApiVersion); err != nil {
		return err
	}
	if m.db, err = fdb.OpenDatabase(m.fdbClusterFile); err != nil {
		return err
	}
	l, err := net.Listen("tcp", m.httpListenAddr)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Print("Stopping fdbmeter...")
				return
			case <-m.statusRefreshInterval.C:
				start := time.Now()
				resp, err := m.db.ReadTransact(func(tr fdb.ReadTransaction) (any, error) {
					return tr.Get(fdb.Key(fdbStatusJsonKey)).Get()
				})
				if err != nil {
					log.Printf("Failed to get status: %v", err)
					m.status.Store(nil)
					m.metrics.notifyStatusTransactFailed(ctx)
					continue
				}
				switch v := resp.(type) {
				case []byte:
					var currentStatus Status
					if err = json.Unmarshal(v, &currentStatus); err != nil {
						log.Printf("Failed to decode JSON status: %v", err)
						m.status.Store(nil)
						m.metrics.notifyStatusDecodeFailed(ctx)
						continue
					}
					m.status.Store(&currentStatus)
					m.metrics.notifyGetStatusLatency(ctx, time.Since(start))
					m.metrics.notifyStatus(ctx, currentStatus)
					log.Print("Refreshed status successfully")
				default:
					log.Printf("Unexpected status response kind: %v", v)
					m.metrics.notifyStatusDecodeFailed(ctx)
				}
			}
		}
	}()
	m.server = &http.Server{
		Handler: m.ServeMux(),
	}
	go func() { _ = m.server.Serve(l) }()
	m.server.RegisterOnShutdown(cancel)
	return nil
}

func (m *FDBMeter) Shutdown(ctx context.Context) error {
	if m.metrics != nil {
		_ = m.metrics.shutdown(ctx)
	}
	if m.server != nil {
		return m.server.Shutdown(ctx)
	}
	return nil
}
