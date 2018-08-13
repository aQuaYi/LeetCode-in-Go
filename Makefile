run:
	@./helper readme
	git commit -am '更新 README.md'
	git push
	git checkout master
	git merge develop
	git push
	git checkout develop
	google-chrome https://github.com/aQuaYi/LeetCode-in-Go#leetcode-%E7%9A%84-go-%E8%A7%A3%E7%AD%94