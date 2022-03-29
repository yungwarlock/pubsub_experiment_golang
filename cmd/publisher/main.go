// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/nerdraven/pubsub_experiment/pkg/protos"
	"github.com/nerdraven/pubsub_experiment/pkg/pubsub"
	"google.golang.org/protobuf/proto"
)

var message string

func init() {
	const (
		defaultMessage = "Pubsub message"
		usage          = "the event name"
	)
	flag.StringVar(&message, "message", defaultMessage, usage)
	flag.StringVar(&message, "n", defaultMessage, usage+" (shorthand)")
}

func main() {
	flag.Parse()
	fmt.Println(message)

	ctx := context.Background()

	publisher, err := pubsub.New("eventtype", "dummy-project")
	if err != nil {
		log.Fatal(err)
	}

	id := fmt.Sprint(rand.Int())

	pb := &protos.Event{Id: id, Name: message}

	pbbytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
	}

	if err := publisher.Publish(ctx, pbbytes, map[string]string{
		"Type": "Event",
	}); err != nil {
		log.Fatal(err)
	}

}
