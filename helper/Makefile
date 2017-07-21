build:
	@go build
	@mv helper ../Helper
	@cd .. ; ./Helper
	
test:	
	go test -v -race -cover -coverprofile cover.out 
	go tool cover -html=cover.out
