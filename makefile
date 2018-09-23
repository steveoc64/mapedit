all: run

test:
	go vet ./...
	golint ./...
	go test ./...

build: test
	cd cmd && make

run: build
	./cmd/mapedit
	
