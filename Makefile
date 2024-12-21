.PHONY: create
create:
	docker buildx create --use --platform=linux/arm64,linux/amd64 --name multi-platform-builder
.PHONY: login
login:
	docker logout ghcr.io
	docker login ghcr.io
.PHONY: build
build:
	docker buildx build --platform linux/arm64,linux/amd64 --push --tag ghcr.io/sharkboy-j/cws:latest --tag ghcr.io/sharkboy-j/cws:0.7 .
#.PHONY: push
#push:
#		docker push ghcr.io/sharkboy-j/cws --all-tags

