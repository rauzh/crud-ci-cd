.PHONY: build
build:
	go build -mod=vendor -o bin/app ./notestorage

.PHONY: docker
docker:
	docker build -t 05420886/ci-ci-2022:latest .
	docker push 05420886/ci-ci-2022:latest

.PHONY: lint
lint:
	golangci-lint -c .golangci.yml run ./notestorage/main.go

