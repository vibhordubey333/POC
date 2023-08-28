package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
)

var (
	projectID   = flag.String("project", "cpl-ddcs-l-sandbox-01", "GCP project-id")
	topicID     = flag.String("topic", "projects/cpl-ddcs-l-sandbox-01/topics/pubsubexample", "Topic ID")
	csvFilePath = flag.String("csv", "./csv-file/movies.csv", "Path to csv file")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	//Initialize a Pub/Sub client
	client, err := pubsub.NewClient(ctx, *projectID, option.WithCredentialsFile("/home/vagrant/Credentials_21june2023.json"))
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	// Open the CSV file.
	file, err := os.Open(*csvFilePath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Create a Pub/Sub topic client.
	topic := client.Topic(*topicID)

	// Create a CSV reader.
	csvReader := csv.NewReader(file)

	// Read and publish each row from the CSV file.
	for {
		record, err := csvReader.Read()
		if err != nil {
			break // End of file
		}

		// Assuming the first column in the CSV contains the message.
		message := &pubsub.Message{
			Data: []byte(record[0]), // You can modify this to fit your CSV structure.
		}

		// Publish the message to the Pub/Sub topic.
		result := topic.Publish(ctx, message)
		_, err = result.Get(ctx)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			fmt.Printf("Published message: %s\n", record[0])
		}
	}
}
