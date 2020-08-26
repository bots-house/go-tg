platform = $(shell uname | tr '[:upper:]' '[:lower:]')

goswagger_version = 0.24.0
goswagger_download_url = $(shell curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/tags/v$(goswagger_version) | jq -r '.assets[] | select(.name | contains("$(platform)_amd64")) | .browser_download_url')

sqlboiler_version = 4.1.2
sqlboiler_download_url = https://api.github.com/repos/volatiletech/sqlboiler/tarball/v$(sqlboiler_version)

golangci_lint_version = 1.27.0

hosts_file=/etc/hosts
localhost_alias=local.birzzha.me
hosts_file_contains_localhost_alias := $(shell grep -q ${localhost_alias} ${hosts_file}; echo $$?)

run: services
	go run main.go -config .env.local

run-server: services
	go run main.go -config .env.local -server 

run-worker: services
	go run main.go -config .env.local -worker

services:
	docker-compose up --no-recreate --detach postgres redis 

run-frontend: frontend-localhost-alias
	@$(info 👇pull latest images of web and admin)
	@docker-compose pull --quiet web admin
	
	@$(info 👟restart containers)
	@docker-compose up --quiet --detach --force-recreate web admin
	
	@$(info 🚀 site: http://${localhost_alias}:80)
	@$(info 🚀 admin: http://${localhost_alias}:8080)

frontend-localhost-alias:
ifeq ($(hosts_file_contains_localhost_alias), 1)
	@$(info ❓host ${localhost_alias} not found in ${hosts_file}, next command will add it)
	echo '127.0.0.1 ${localhost_alias}' | sudo tee -a /etc/hosts
endif

psql:
	docker-compose exec postgres psql -U brz

redis-cli:
	docker-compose exec redis redis-cli

lint: lint-go lint-swagger

lint-go: .bin/golangci-lint
		.bin/golangci-lint run --config .golangci.yml

lint-swagger: .bin/swagger
	.bin/swagger validate ./api/swagger.yml --stop-on-error

generate:  generate-dal generate-api

generate-dal: .bin/sqlboiler .bin/sqlboiler-psql
	.bin/sqlboiler .bin/sqlboiler-psql

generate-api: .bin/swagger
	rm -r api/gen && mkdir api/gen
	cd api && mkdir -p gen && ../.bin/swagger generate server --quiet --spec swagger.yml --exclude-main --keep-spec-order --target=gen --principal=authz.Identity

.bin/sqlboiler .bin/sqlboiler-psql:
	mkdir -p .bin
	curl -o .bin/sqlboiler.tar.gz -L $(sqlboiler_download_url)
	tar -xzf .bin/sqlboiler.tar.gz --directory .bin && rm .bin/sqlboiler.tar.gz
	mv .bin/volatiletech-sqlboiler-* .bin/sqlboiler-src
	cd .bin/sqlboiler-src && go build -o ../sqlboiler
	cd .bin/sqlboiler-src/drivers/sqlboiler-psql && go build -o ${CURDIR}/.bin/sqlboiler-psql
	rm -r .bin/sqlboiler-src

.bin/swagger:
	mkdir -p .bin
	curl -o $@ -L $(goswagger_download_url)
	chmod +x $@

.bin/golangci-lint:
	mkdir -p .bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b .bin v$(golangci_lint_version)
