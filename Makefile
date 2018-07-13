.PHONY: bombard_server run

bombard_server:
	sh bombard.sh

run:
	@go run main.go
