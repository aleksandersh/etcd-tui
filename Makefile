demo-etcd-docker-up:
	docker compose -f ./demo/docker-compose.yml up -d

demo-etcd-docker-down:
	docker compose -f ./demo/docker-compose.yml down

demo-connect: build
	./etcd-tui localhost:2379

build:
	go build
