package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/masih/fdbmeter"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	httpListenAddr := flag.String("httpListenAddr", "0.0.0.0:40080", "The bind address of fdbmeter HTTP server.")
	fdbApiVersion := flag.Int("fdbApiVersion", 730, "The FoundationDB API version.")
	fdbClusterFile := flag.String("fdbClusterFile", "", "Path to the FoundationDB cluster file.")
	statusRefreshInterval := flag.Duration("statusRefreshInterval", 10*time.Second, "The interval at which to refresh the FoundationDB status.")
	commonAttributes := flag.String("commonAttributes", "", "The common attributes to apply to all metrics specified as comma separated key=value.")
	flag.Parse()

	var commonAttrs []attribute.KeyValue
	if *commonAttributes != "" {
		attrString := strings.Split(*commonAttributes, ",")
		for _, attr := range attrString {
			kv := strings.Split(attr, "=")
			if len(kv) != 2 {
				log.Fatalf("Invalid common attributes: '%s'. Expected exactly one key and one value per comma separated attribute.", attr)
			}
			commonAttrs = append(commonAttrs, attribute.String(kv[0], kv[1]))
		}
	}
	meter, err := fdbmeter.New(
		fdbmeter.WithHttpListenAddr(*httpListenAddr),
		fdbmeter.WithFdbApiVersion(*fdbApiVersion),
		fdbmeter.WithFdbClusterFile(*fdbClusterFile),
		fdbmeter.WithStatusRefreshInterval(*statusRefreshInterval),
		fdbmeter.WithCommonAttributes(commonAttrs...),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	if err = meter.Start(ctx); err != nil {
		log.Fatal("Failed to start fdbmeter ", err)
	}
	log.Print("Started fdbmeter with listen address: ", *httpListenAddr)
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, os.Interrupt)

	<-sch
	cancel()
	log.Print("Terminating...")
	if err := meter.Shutdown(ctx); err != nil {
		log.Printf("Failure occurred while shutting down server: %v", err)
	} else {
		log.Print("Shut down server successfully.")
	}
}
