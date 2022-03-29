
GO = go

BUILD_DIR = build

DOCKER_BIN = gcr.io/$(PROJECT_ID)/subscriber:latest

PROTO_ROOT = .
PROTO_OUT = ./cmd/protos/
PROTO_OPTS = --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

build:
	mkdir $(BUILD_DIR)

all: publisher subscriber

protos:
	mkdir $(PROTO_OUT)
	protoc -I$(PROTO_ROOT) $(PROTO_OPTS)  --go-grpc_out=$(PROTO_OUT) --go_out=$(PROTO_OUT)  $(PROTO_ROOT)/event.proto

publisher: build
	$(GO) build -o $(BUILD_DIR)/publisher ./cmd/publisher

subscriber: build
	$(GO) build -o $(BUILD_DIR)/subscriber ./cmd/subscriber

deploy_subscriber: subscriber
	gcloud builds submit --tag $(DOCKER_BIN) .
	gcloud run deploy pubsub-subscriber  --allow-unauthenticated --image=$(DOCKER_BIN)

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(PROTO_OUT)
	