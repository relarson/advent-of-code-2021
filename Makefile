folder = day$(day)

build:
	go build -o bin/$(folder) cmd/$(folder)/main.go

run:
	go run cmd/$(folder)/main.go