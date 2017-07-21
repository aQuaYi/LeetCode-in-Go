test:	
	go test -v -race -cover -coverprofile cover.out 
	go tool cover -html=cover.out

build:
	@go build
	@mv helper ../Aida