build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/save_bar handler/save_bar/main.go