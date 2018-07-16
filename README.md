# throttle-api-using-kafka-poc

A proof of concept for throttling an API endpoint using Kafka and Go.

## !Main TODOs!
  
* Use offsets so that the Kafka Consumer knows where to start processing stored requests if the server ever stops/crashes.

## Prerequisites
* [Go 1.10.*](https://golang.org/)
* [Git](https://git-scm.com/)
* [dep (Go package manager)](https://golang.github.io/dep/)
* [Apache Kafka](https://kafka.apache.org/)
* [Apache Kafka C/C++ library](https://github.com/edenhill/librdkafka)

**Quick install for MacOS**

`brew install go git dep kafka java8 librdkafka`

**Start Apache Kafka (using MacOS and `brew`)**
```
brew services start zookeeper
brew services start kafka
```

## Installation

```
go get -u -v github.com/angelospanag/throttle-api-using-kafka-poc
dep ensure -v
```

## Configuration

Place a config.toml file with the following content at the root of the project:
```toml
[kafka]
servers = ['localhost:9092']
topic = 'throttled_topic'

[consumption]
number_of_messages = 5
time_period_seconds = 5
```

## Running

`go run main.go` or `make run`

Port used: `8080`

## Architecture and usage

The purpose of this project is to showcase a throttling functionality for an exposed API endpoint. After initiating the project server and having a Kafka instance running you can try it yourself.

### Kafka Consumer and Producer
The server initiates a Kafka Consumer and Producer as soon as it begins executing.

### Bombard your server
Start by bombarding your server at the root URL endpoint (`/`) with HTTP POST requests by executing `make bombard_server`. This will execute a script that will send 100 requests that contain a JSON body with content `{"number": NUMBER_OF_REQUEST}`. The called endpoint will return an HTTP OK to a user after posting a JSON message but at the same time it will forward this message asynchronously to a Kafka topic.

### Throttling
A Kafka polling goroutine also begins executing periodically as soon as the server initiates. It uses the previously initialised Kafka Consumer for polling new messages existing in the topic that was used above. It fetches them and prints them out in the order they were received from your previous bombardment and at a rate of 5 messages every 5 seconds.
