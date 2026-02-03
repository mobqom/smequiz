include .env

build:
	@go build -o ./dist/app 
run: build
	@./dist/app	
test:
	@echo export HOST=HOST
	@echo export PORT=PORT
	@go clean -testcache
	@go test -v ./test/connection_test.go
test-race: 
	@echo export HOST=HOST
	@echo export PORT=PORT
	@go clean -testcache
	@go test -race -v ./test
