.PHONY: docker-up
docker-up:
	docker-compose --env-file ./scripts/dbcredentials.env up -d
	./scripts/setup.sh


.PHONY: run
run:
	go run main.go