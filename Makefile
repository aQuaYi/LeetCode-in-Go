run:
	./helper
	git commit -am '更新README.md'
	git push
	git checkout master
	git merge develop
	git push
	git checkout develop