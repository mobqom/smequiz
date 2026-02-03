build:
	@go build -o ./dist/app 
run: build
	@./dist/app	
test:
	@go clean -testcache
	@go test ./test -run TestConnection -v

test-race: 
	@go clean -testcache
	@go test ./test -run TestConnection -race -v

