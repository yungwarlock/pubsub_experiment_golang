package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type Publisher interface {
	Publish(context.Context, []byte) error
}

type PublisherImpl struct {
	topic *pubsub.Topic
}

func New(topicId string, projectId string) (*PublisherImpl, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}
	topic := client.Topic(topicId)
	ok, err := topic.Exists(ctx)
	if err != nil || !ok {
		return nil, err
	}
	return &PublisherImpl{topic: topic}, nil
}

func (p *PublisherImpl) Publish(ctx context.Context, message []byte, attributes map[string]string) error {
	pubsubmesg := pubsub.Message{
		Data:       message,
		Attributes: attributes,
	}

	res := p.topic.Publish(ctx, &pubsubmesg)

	// Verify if the message has been published successfully
	_, err := res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
