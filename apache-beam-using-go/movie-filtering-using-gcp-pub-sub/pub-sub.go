package main

import (
	"flag"
)

var (
	projectID   = flag.String("project", "cpl-ddcs-l-sandbox-01", "GCP project-id")
	topicID     = flag.String("topic", "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", "Topic ID")
	csvFilePath = flag.String("csv", "", "Path to csv file")
)
