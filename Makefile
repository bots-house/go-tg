platform = $(shell uname | tr '[:upper:]' '[:lower:]')

goswagger_version = 0.24.0
goswagger_download_url = $(shell curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/tags/v$(goswagger_version) | jq -r '.assets[] | select(.name | contains("$(platform)_amd64")) | .browser_download_url')

sqlboiler_version = 4.1.2
sqlboiler_download_url = https://api.github.com/repos/volatiletech/sqlboiler/tarball/v$(sqlboiler_version)

run:
	go run main.go

generate:  generate-dal generate-api

generate-dal: .bin/sqlboiler .bin/sqlboiler-psql
	.bin/sqlboiler .bin/sqlboiler-psql

generate-api: .bin/swagger
	cd api && mkdir -p gen && swagger generate server --quiet --spec swagger.yml --exclude-main --keep-spec-order --target=gen --principal=authz.Identity

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
