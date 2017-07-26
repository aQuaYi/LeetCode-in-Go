build:
	@go build
	@mv helper ../Helper
	@cd .. ; ./Helper
	@git commit -am '更新helper助手和README.md'
	
test:	
	go test -v -race -cover -coverprofile cover.out 
	go tool cover -html=cover.out
