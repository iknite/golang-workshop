.PHONY: slides install run test clean
DEFAULT: install

vendor/:
	mkdir vendor

ENV/GOPATH: vendor/
	@echo "$(realpath vendor)" > ENV/GOPATH

bin/:
	mkdir bin

node_modules/: 
	npm install

slides: node_modules/
	./node_modules/.bin/reveal-md README.md --disableAutoOpen --theme moon

install: vendor/
	@go get $(shell cat dependencies.txt)

run: ENV/GOPATH
	@go run src/pitytweet/main.go --name $$NAME --key $$KEY --secret $$SECRET

test: ENV/GOPATH
	@go test src/pitytweet/main.go

build: bin/ ENV/GOPATH
	@go build -o bin/pitytweet src/pitytweet/main.go

clean:
	@rm -rf vendor/
	@rm -rf node_modules/

