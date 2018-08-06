buildlocal:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/statsservice	\
		proto/statsservice.proto
build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/Nanixel/FirstDownMicro/statsservice	\
		proto/statsservice.proto
	docker build -t statsservice .
runlocal:
	go run *.go --server_address :50052 --registry mdns --broker nats --broker_address :4222
run:
	docker run --net="host" \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e MICRO_REGISTRY=mdns \
		-e MICRO_BROKER=nats \
		-e MICRO_BROKER_ADDRESS=:4222 statsservice