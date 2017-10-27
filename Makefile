run:
	@./helper readme
	git commit -am '更新 README.md'
	git push
	git checkout master
	git merge develop
	git push
	git checkout develop