build:
	go build -o ./.bin/raycast cmd/raycast.go

run: build
	./.bin/raycast
