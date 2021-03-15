REPO=marcoamador
NAME=pipeline
VERSION=0.1.1

all: docker
clean: docker-clean
build: compile
test: godog

run:
	go run cmd/pipeline/main.go -c config/pipeline.yaml

compile:
	go build -o /dev/null -ldflags "-s -w" ./cmd/pipeline
	# compiling without binary output (remove -o /dev/null if you want to generate a binary)

deps:
	go mod download && go mod tidy

godog:
	(cd test && godog)

docker:
	docker build -f build/Dockerfile -t $(REPO)/$(NAME):$(VERSION) .

docker-push:
	docker push $(REPO)/$(NAME):$(VERSION)

docker-clean:
	docker rmi $(REPO)/$(NAME):$(VERSION)

docker-run:
	docker run -f build/Dockerfile --rm --name $(NAME) $(REPO)/$(NAME):$(VERSION) --squash
