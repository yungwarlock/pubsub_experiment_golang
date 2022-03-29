package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nerdraven/pubsub_experiment/pkg/protos"
	"github.com/nerdraven/pubsub_experiment/pkg/pubsub"
	"google.golang.org/protobuf/proto"
)

var (
	port = os.Getenv("PORT")
)

func publish(w http.ResponseWriter, r *http.Request) {
	var m pubsub.PubSubMessage

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &m); err != nil {
		log.Printf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var data protos.Event
	proto.Unmarshal(m.Message.Data, &data)

	fmt.Printf("Event of id %s has been sent to pubsub", data.Id)

	fmt.Fprintf(w, "Recieved event %s", data.Name)

}

func main() {
	http.HandleFunc("/publish", publish)
	http.ListenAndServe(":"+port, nil)
}
