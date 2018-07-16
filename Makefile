.PHONY: delete_topic_data bombard_server run 

topic=throttled_topic
zookeeper_url=localhost:2181

# This will delete all data currently existing in a Kafka topic, use variables from above
delete_topic_data:
	kafka-topics --zookeeper $(zookeeper_url) --delete --topic $(topic)

bombard_server:
	sh bombard.sh

run:
	@go run main.go

