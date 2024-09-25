help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    docker v                : show docker version"
	@echo "    linux                   : install docker for linux"
	@echo "    macos                   : install docker for macos"
	@echo "    env                     : install base project env"
	@echo "    mongo                   : install mongodb"
	@echo "    mod                     : go set env and go install package"
	@echo "    run                     : run EchoFrame"
	@echo "    clean                   : clean local mod cache"

PWD := $(shell pwd)

# app项目名
AppName := app

# Go 命令
GO := go

docker v:
	docker version


linux:
	chmod +x $(PWD)/scripts/docker/install_docker_linux.sh
	scripts/docker/install_docker_linux.sh

macos:
	chmod +x $(PWD)/scripts/docker/install_docker_macos.sh
	scripts/docker/install_docker_macos.sh

env:
	chmod +x $(PWD)/scripts/env/deploy_env.sh
	scripts/env/deploy_env.sh

mongo:
	chmod +x $(PWD)/scripts/env/deploy_mongo.sh
	scripts/env/deploy_mongo.sh

mod:
	$(GO) mod tidy

run:
	$(GO) run cmd/app/main.go

clean:
	$(GO) clean -modcache


