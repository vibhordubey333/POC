package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

func main() {
	pflag.String("subscription", "example-subscription", "name of subscriber")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	ctx := context.Background()
	proj := "cpl-ddcs-l-sandbox-01"
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	sub := viper.GetString("subscriber") // retrieve values from viper instead of pflag
	if len(sub) <= 0 {
		sub = "PubSub-Demo-sub-2"
	}
	fmt.Println("subscriber:", sub)
	t := createTopicIfNotExists(client, "pubsubexample")
	// Create a new subscription.
	if err := create(client, sub, t); err != nil {
		log.Fatal("Create: ", err)
	}
	// Pull messages via the subscription.
	if err := pullMsgs(client, sub, t, false); err != nil {
		log.Fatal(err)
	}
}

func createTopicIfNotExists(c *pubsub.Client, topic string) *pubsub.Topic {
	ctx := context.Background()
	t := c.Topic(topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		return t
	}
	t, err = c.CreateTopic(ctx, topic)
	if err != nil {
		log.Fatalf("Failed to create the topic: %v", err)
	}
	log.Println("CreateTopicIfNotExists: ", t.String())
	return t
}

func create(client *pubsub.Client, name string, topic *pubsub.Topic) error {
	ctx := context.Background()
	// [START pubsub_create_pull_subscription]
	sub, err := client.CreateSubscription(ctx, name, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		log.Fatalln("Inside create:\n", err)
		//return err
	}
	fmt.Printf("Created subscription: %v\n", sub)
	// [END pubsub_create_pull_subscription]
	return nil
}

func pullMsgs(client *pubsub.Client, name string, topic *pubsub.Topic, testPublish bool) error {
	ctx := context.Background()
	fmt.Println("Inside pullMsgs")
	if testPublish {
		// Publish 10 messages on the topic.
		var results []*pubsub.PublishResult
		for i := 0; i < 10; i++ {
			res := topic.Publish(ctx, &pubsub.Message{
				Data: []byte(fmt.Sprintf("hello world #%d", i)),
			})
			results = append(results, res)
		}

		// Check that all messages were published.
		for _, r := range results {
			_, err := r.Get(ctx)
			if err != nil {
				return err
			}
		}
	}
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(name)
	cctx, cancel := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		mu.Lock()
		defer mu.Unlock()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return err
	}
	return nil
}
