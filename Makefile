.PHONY: build
build:
	go build
	docker build -t ghcr.io/rebell81/cws:0.4 -t ghcr.io/rebell81/cws:latest --platform linux/amd64 --no-cache .
.PHONY: push
push:
	docker push ghcr.io/rebell81/cws --all-tags

.PHONY: login
login:
	docker logout ghcr.io
	docker login ghcr.io

.PHONY: build-push
build-push: build push