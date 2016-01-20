.PHONY: slides install gopath run test clean
DEFAULT: install

vendor/:
	mkdir vendor

bin/:
	mkdir bin

node_modules/: 
	npm install


slides: node_modules/
	./node_modules/.bin/reveal-md README.md --disableAutoOpen --theme moon

install: vendor/
	@eval $(shell make gopath) && go get $(shell cat dependencies.txt)

gopath: vendor/
	@# eval this output
	@echo "export GOPATH=$(realpath vendor)"

run: 
	@eval $(shell make gopath) && go run src/pitytweet/main.go --name $(shell cat env/NAME)

test:
	@eval $(shell make gopath) && go test src/pitytweet/main.go

build: bin/
	@eval $(shell make gopath) && go build -o bin/pitytweet src/pitytweet/main.go

clean:
	@rm -rf vendor/
	@rm -rf node_modules/

