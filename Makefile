PKG := github.com/deuanz/golang-with-heroku
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GOLINT?=		go run golang.org/x/lint/golint

docker/local/db/up:
	@echo "============= starting db locally ============="s
	go mod tidy
	docker-compose -f resources/docker/docker-compose.yaml up postgres

docker/local/migrate:
	docker-compose -f resources/docker/docker-compose.yaml up -d postgres
	go mod tidy
	docker-compose -f resources/docker/docker-compose.yaml -f resources/docker/docker-compose-migrate.yaml up --build golang-with-heroku

docker/local/migrate-api-secret:
	docker-compose -f resources/docker/docker-compose.yaml up -d postgres
	go mod tidy
	RSA="$(rsa)" docker-compose -f resources/docker/docker-compose.yaml -f resources/docker/docker-compose-migrate-api-secret.yaml up --build golang-with-heroku

docker/local/seed:
	docker-compose -f resources/docker/docker-compose.yaml up -d postgres
	go mod tidy
	docker-compose -f resources/docker/docker-compose.yaml -f resources/docker/docker-compose-seed.yaml up --build golang-with-heroku

docker/local/up:
	@echo "============= starting locally ============="
	go mod tidy
	docker-compose -f resources/docker/docker-compose.yaml up --build

lint: ## Lint the files
	$(GOLINT) -set_exit_status ${PKG_LIST}

down:
	@echo "============= downing core service ============="
	docker-compose -f resources/docker/docker-compose.yaml down

test:
	@go test -v ${PKG_LIST}

migrate:
	migrate create -ext sql -dir app/migrations $(name)

format:
	go fmt ./app/...

#mock/all:


mock/repo:
	mockgen \
		-source=./app/domain/repos/$(m)/main.go \
		-destination=./app/domain/repos/$(m)/mocks/$(m).go \
		-package $(m)repomocks \
		-mock_names Repo=Mocks

mock/usecase:
	mockgen \
		-source=./app/domain/usecases/$(m)/main.go \
		-destination=./app/domain/usecases/$(m)/mocks/$(m).go \
		-package $(m)usecasemocks \
		-mock_names UseCase=Mocks

mock/interfaces/adapters:
	mockgen \
		-source=./app/domain/interfaces/adapters/${file_name} \
		-package adaptermocks \
		-destination=./app/domain/interfaces/adapters/mocks/${file_name}

mock/interfaces:
	mockgen \
		-source=./app/domain/interfaces/${module}/main.go \
		-package ${module}mocks \
		-destination=./app/domain/interfaces/${module}/mocks/main.go

seed-scb:
	curl -X POST http://0.0.0.0:8001/dev/seed -d '{"code": "SCB", "name": "SCB", "legalName": "Siam Commercial Bank"}'
