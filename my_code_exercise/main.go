package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/honeycombio/libhoney-go"
)



func quit(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", v...)
	os.Exit(1)
}

func main() {
	var (
		apiKey string
		dataset string
		markerType string
		message string
		del bool
	)

	flag.StringVar(&apiKey, "k", "", "honeycomb api key")
	flag.StringVar(&dataset, "d", "", "honeycomb dataset name")
	flag.StringVar(&markerType, "t", "test", "marker type")
	flag.StringVar(&message, "m", "I am a marker", "marker message")
	flag.BoolVar(&del, "x", false, "delete the marker that you just created")
	flag.Parse()

	if apiKey == "" {
		quit("honeycomb api key is required")
	}
	if dataset == "" {
		quit("honeycomb dataset name is required")
	}

	initErr := libhoney.Init(libhoney.Config{
		APIKey: apiKey,
		Dataset: dataset,
	})
	if initErr != nil {
		quit("libhoney.Init error: %v", initErr)
	}

	marker := libhoney.Marker{
		Type:      markerType,
		Message:   message,
	}

	fmt.Printf("sending marker:\n\t%+v\n", marker)
	id, err := marker.Send()
	if err != nil {
		quit("send marker failed: %v", err)
	}

	fmt.Printf("marker id: %s\n", id)
	<-time.After(time.Second)

	if !del {
		os.Exit(0)
	}

	if err := libhoney.DeleteMarker(id); err != nil {
		quit("delete marker failed: %v", err)
	}
}
