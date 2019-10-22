DOCKER_COMPOSE_VERSION=1.24.0
NAMESPACE=sr2020
SERVICE := platform
IMAGE := $(or ${image},${image},eva-gateway)
GIT_TAG := $(shell git tag -l --points-at HEAD | cut -d "v" -f 2)
TAG := :$(or ${tag},${tag},$(or ${GIT_TAG},${GIT_TAG},latest))
ENV := $(or ${env},${env},local)
cest := $(or ${cest},${cest},)

current_dir = $(shell pwd)

build:
	docker build -t ${NAMESPACE}/${IMAGE}${TAG} .

push:
	docker push ${NAMESPACE}/${IMAGE}

deploy:
	{ \
	sshpass -p $(password) ssh -o StrictHostKeyChecking=no deploy@$(server) "cd /var/services/$(SERVICE) ;\
	docker-compose pull gateway-app ;\
	docker-compose up -d --no-deps gateway-app" ;\
	}

deploy-app:
	cd ansible && ansible-playbook -i inventories/production -u deploy deploy-app.yml --extra-vars ansible_ssh_pass=$(password)

deploy-local:
	docker-compose rm -fs app
	docker-compose up --no-deps app

up:
	docker-compose up -d

dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up

down:
	docker-compose down

reload:
	make down
	make up

restart:
	docker-compose down -v
	docker-compose up -d

install:
	cp .env.example .env

install-docker-compose:
	curl -L https://github.com/docker/compose/releases/download/$(DOCKER_COMPOSE_VERSION)/docker-compose-Linux-x86_64 > /tmp/docker-compose
	chmod +x /tmp/docker-compose
	sudo mv /tmp/docker-compose /usr/local/bin/docker-compose
	docker-compose -v

test:
	docker run -v $(current_dir)/tests:/project --net host codeception/codeception run $(ENV) $(cest)

load:
	docker run -v $(current_dir)/tests/loadtest:/var/loadtest --net host --entrypoint /usr/local/bin/yandex-tank -it direvius/yandex-tank -c production.yaml

test-dev:
	make build
	make up
	make test

database-dump-update:
	wget -O database/auth.sql https://raw.githubusercontent.com/sr-2020/eva-auth/master/docker/mysql/dump.sql
	wget -O database/position.sql https://raw.githubusercontent.com/sr-2020/eva-position/master/docker/mysql/dump.sql
