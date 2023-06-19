package fdbmeter

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
)

type (
	Option  func(*options) error
	options struct {
		httpListenAddr        string
		fdbApiVersion         int
		fdbClusterFile        string
		statusRefreshInterval *time.Ticker
		commonAttributes      []attribute.KeyValue
	}
)

func newOptions(o ...Option) (*options, error) {
	opts := options{
		httpListenAddr:        "0.0.0.0:40080",
		fdbApiVersion:         710,
		statusRefreshInterval: time.NewTicker(10 * time.Second),
	}
	for _, apply := range o {
		if err := apply(&opts); err != nil {
			return nil, err
		}
	}
	return &opts, nil
}

func WithHttpListenAddr(a string) Option {
	return func(o *options) error {
		o.httpListenAddr = a
		return nil
	}
}

func WithFdbApiVersion(v int) Option {
	return func(o *options) error {
		o.fdbApiVersion = v
		return nil
	}
}

func WithFdbClusterFile(cf string) Option {
	return func(o *options) error {
		o.fdbClusterFile = cf
		return nil
	}
}

func WithStatusRefreshInterval(d time.Duration) Option {
	return func(o *options) error {
		o.statusRefreshInterval = time.NewTicker(d)
		return nil
	}
}

func WithCommonAttributes(a ...attribute.KeyValue) Option {
	return func(o *options) error {
		o.commonAttributes = a
		return nil
	}
}
