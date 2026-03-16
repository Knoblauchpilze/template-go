
# https://stackoverflow.com/questions/34712972/in-a-makefile-how-can-i-fetch-and-assign-a-git-commit-hash-to-a-variable
GIT_COMMIT_HASH=$(shell git rev-parse --short HEAD)
SERVER_PORT=60004
DATABASE_HOST=127.0.0.1
DATABASE_PASSWORD="comes-from-the-environment"

template-go-build:
	docker build \
		--build-arg GIT_COMMIT_HASH=${GIT_COMMIT_HASH} \
		--tag totocorpsoftwareinc/template-go:${GIT_COMMIT_HASH} \
		-f build/template-go/Dockerfile \
		.

template-go-run:
	docker run \
		-e ENV_SERVER_PORT=${SERVER_PORT} \
		-e ENV_DATABASE_HOST=${DATABASE_HOST} \
		-e ENV_DATABASE_PASSWORD=${DATABASE_PASSWORD} \
		-p 1234:1234 \
		--network=host \
		totocorpsoftwareinc/template-go:${GIT_COMMIT_HASH}
