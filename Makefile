#PROJECT_ROOT_DIR hold the absolute path to this project root dir.
PROJECT_ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#
# Test Related Commands
#

.PHONY: build-test
build-test:
	@docker build --target test -f ./build/docker/dev/go/Dockerfile -t meli-test:latest .

.PHONY: test
test:
	@go test -v ./... -coverpkg="./internal/...,./pkg/...,./cmd/..." -cover -coverprofile=./coverage.txt -covermode=atomic -gcflags="all=-N -l"

.PHONY: coverage
coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@go tool cover -func coverage.txt -o coverage.func.txt

.PHONY: docker-test
docker-test:
	@docker run --rm -v ${PROJECT_ROOT_DIR}:/go/src/gitlab.com/meli meli-test:latest make test

.PHONY: docker-coverage
docker-coverage:
	@docker run --rm -v ${PROJECT_ROOT_DIR}:/go/src/gitlab.com/meli meli-test:latest make coverage

#
# Dev Reletad Commands
#
.PHONY: build-dev
build-dev:
	@docker-compose build

.PHONY: dev
dev: build-dev
	@docker-compose up go

.PHONY: stop down
down: stop
stop:
	@docker-compose stop

#Container must be running
.PHONY: logs
logs:
	@docker logs -f meli-go

#
# Remote debug related Commands
#

.PHONY: delve
delve:
	@dlv debug ./cmd/api --listen=:33333  --headless=true --api-version=2 --output=./tmp/__debug_bin --log

.PHONY: debug
debug: githooks build-dev
	@docker-compose up debug

#swagger
.PHONY: docker-swagger
docker-swagger:
	@docker exec -it meli-test swag init --generalInfo cmd/api/main.go --output ./api/v1/docs
