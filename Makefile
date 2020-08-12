push:
	git pull origin master
	git add -A
	git commit -S
	git push origin master

.PHONY: push
