demo-etcd-docker-up:
	docker compose -f ./demo/docker-compose.yml up -d

demo-etcd-docker-down:
	docker compose -f ./demo/docker-compose.yml down

demo-connect: build
	./etcd-tui localhost:2379

build:
	go build

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 .run/build.sh

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 .run/build.sh

build-linux-amd64:
	GOOS=linux GOARCH=amd64 .run/build.sh

build-linux-arm64:
	GOOS=linux GOARCH=arm64 .run/build.sh

build-windows-amd64:
	GOOS=windows GOARCH=amd64 .run/build.sh exe

build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64
