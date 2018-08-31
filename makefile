build:
	go get ./...
	go build -o nsq-transport

deploy-local:
	make build
	mv nsq-transport /usr/local/bin/
	chmod +x /usr/local/bin/nsq-transport
