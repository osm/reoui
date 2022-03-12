.PHONY: all gql frontend clean release release-armv6

all:
	go build

gql:
	go run github.com/99designs/gqlgen generate

frontend:
	cd frontend && yarn && yarn build

clean:
	rm -f reoui
	rm -f frontend/dist/*
	rm -rf frontend/node_modules

release:
	cd frontend && yarn && yarn build
	go build

release-armv6:
	cd frontend && yarn && yarn build
	GOOS=linux GOARCH=arm GOARM=6 go build
