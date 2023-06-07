build:
	go build -o tmp/allfather cmd/app/main.go

run: build
	./tmp/allfather