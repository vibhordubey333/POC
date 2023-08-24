package main

import (
	"context"
	"flag"
	_ "fmt"
	"log"

	_ "cloud.google.com/go/pubsub"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	_ "google.golang.org/api/option"
)

var (
	project1 = flag.String("project1", "cpl-ddcs-l-sandbox-01", "Google Cloud Project ID")
	input    = flag.String("input", "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", "Input Pub/Sub topic")
	output   = flag.String("output", "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", "Output Pub/Sub topic")
	topic    = flag.String("topic", "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", "Topic Pub/Sub ")
)

func main() {
	flag.Parse()
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	// Read from the input Pub/Sub topic.
	messages := pubsubio.Read(s, *project1, "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", &pubsubio.ReadOptions{
		Subscription:   "projects/cpl-ddcs-l-sandbox-01/subscriptions/pub-sub",
		WithAttributes: false,
	})
	// Process the messages (you can add your custom processing logic here).
	// For example, you can use beam.ParDo to apply a custom function to each message.

	// Write the processed messages to the output Pub/Sub topic.
	pubsubio.Write(s, *project1, "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", messages)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

/*
To run:
go run main.go --runner=dataflow --project=cpl-ddcs-l-sandbox-01 --region=us-central1 --staging_location=gs://data-flow-bucket-vd/kk/staging --sdk_container_image=apache/beam_go_sdk:2.49.0 --output=gs://data-flow-bucket-vd/kk/output --subnetwork=https://www.googleapis.com/compute/v1/projects/cpl-ddcs-l-sandbox-01/regions/us-central1/subnetworks/dataflow-subnet-kk
*/
