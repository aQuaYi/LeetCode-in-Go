run:
	@./helper readme
	git commit -am '更新README.md'
	git checkout master
	git merge develop
	git push
	git checkout develop
	git push