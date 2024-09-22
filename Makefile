help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    docker v                : show docker version"
	@echo "    install docker linux    : install docker for linux"
	@echo "    install docker macos    : install docker for macos"
	@echo "    env                     : install base project env"
	@echo "    mod                     : go set env and go install package"
	@echo "    clean                   : clean this mod file"

# docker
DOCKER := docker

docker v:
	${DOCKER} version


install docker linux:
	chmod +x scripts/docker/install_docker_linux.sh
	scripts/docker/install_docker_linux.sh

install docker macos:
	chmod +x scripts/docker/install_docker_macos.sh
	scripts/docker/install_docker_macos.sh

env:
	chmod +x scripts/env/deploy_env.sh
	scripts/env/deploy_env.sh





