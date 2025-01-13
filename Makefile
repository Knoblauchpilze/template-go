
# https://stackoverflow.com/questions/34712972/in-a-makefile-how-can-i-fetch-and-assign-a-git-commit-hash-to-a-variable
GIT_COMMIT_HASH=$(shell git rev-parse --short HEAD)

template-go-build:
	docker build \
		--build-arg GIT_COMMIT_HASH=${GIT_COMMIT_HASH} \
		--tag totocorpsoftwareinc/template-go:${GIT_COMMIT_HASH} \
		-f build/template-go/Dockerfile \
		.