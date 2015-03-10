USER = pslattery
IMAGE = hello_world
TAG = 1.0.0
TIMEOUT = 2

all: clean build images tag-latest run logs test kill push
kill: stop rm

.PHONY: clean

clean:
	-@docker rmi $(USER)/$(IMAGE):$(TAG)
	-@docker rmi $(USER)/$(IMAGE):latest

cleanimages:
	-docker images -q --filter "dangling=true" | xargs docker rmi

build: Dockerfile
	@go build -o ./config/$(IMAGE) $(IMAGE).go
	@docker \
		build \
		--tag=$(USER)/$(IMAGE):$(TAG) .

tag-latest:
	@docker \
		tag -f \
		$(USER)/$(IMAGE):$(TAG) $(USER)/$(IMAGE):latest

images:
	@docker images | head -n 1 && docker images | grep "$(IMAGE)"

run:
	@docker \
		run \
		--detach \
		--hostname=$(IMAGE) \
		--name=$(IMAGE) \
		--publish="8080:8080" \
		$(USER)/$(IMAGE):$(TAG)

logs:
	-@timeout $(TIMEOUT) docker logs -f $(IMAGE)

test:
	@curl -i -s http://localhost:8080

stop:
	@docker \
		stop $(IMAGE)

rm:
	@docker \
		rm $(IMAGE)

push:
	@docker \
		push \
		$(USER)/$(IMAGE)
