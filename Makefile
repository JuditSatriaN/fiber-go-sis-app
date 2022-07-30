.PHONY: docker-up
docker-up:
	docker-compose --env-file dbcredentials.env up -d
	./utils/schemes/setup.sh


.PHONY: run
run:
	go run main.go