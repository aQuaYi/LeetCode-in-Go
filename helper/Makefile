build:
	@go build
	@mv helper ../Helper
	@cd .. ; ./Helper
	git commit -am '更新helper助手和README.md'
	git push
	git checkout master
	git merge develop
	git push
	git checkout develop
	
test:	
	go test -v -race -cover -coverprofile cover.out 
	go tool cover -html=cover.out
