all: build

build:
	cd cmd && make

run: build
	./cmd/mapedit
	
