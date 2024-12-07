gengo:
	pkl-gen-go pkl/ServiceConfig.pkl

build:
	go build -o dist/sercon main.go
