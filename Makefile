build:
	go build -o ./.bin/game cmd/game.go

run: build
	./.bin/game
