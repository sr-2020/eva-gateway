DOCKER_SSH_KEY := $(shell cat ~/.ssh/id_rsa.pub)
NODE_IP := 172.55.0.10
NODE_NAME := platform-test
NETWORK := platform
SUBNET := 172.55.0.0/16
USER := $(or ${user},${user},root)
ENV := $(or ${env},${env},local)
INVENTORY := inventories/$(ENV)

.PHONY : test

network-create:
	docker network create --subnet=$(SUBNET) $(NETWORK) || true

network-remove:
	docker network rm $(NETWORK) || true

server:
	make network-create
	docker run --net $(NETWORK) --ip $(NODE_IP) --name $(NODE_NAME) -itd -P -v /var/run/docker.sock:/var/run/docker.sock -e SSH_KEY="$(DOCKER_SSH_KEY)" gurkalov/ubuntu-ssh:bionic
	sleep 1
	ssh-keyscan -t rsa -H $(NODE_IP) >> ~/.ssh/known_hosts

down:
	docker rm -f $(NODE_NAME) || true
	ssh-keygen -f ~/.ssh/known_hosts -R $(NODE_IP)
	make network-remove

setup:
	ansible-playbook -i $(INVENTORY) -u $(USER) bootstrap.yml
	ansible-playbook -i $(INVENTORY) -u $(USER) server.yml

deploy:
	ansible-playbook -i $(INVENTORY) -u $(USER) deploy.yml

reboot:
	make down
	make server
	make setup

stage:
	ansible-playbook terraform.yml -i $(INVENTORY)
	sleep 120
	make setup user=$(USER) env=$(ENV)
	make deploy user=$(USER) env=$(ENV)
	sleep 500
	ansible-playbook -i $(INVENTORY) -u $(USER) restore.yml
	sleep 60
	ansible-playbook -i $(INVENTORY) -u $(USER) deploy-app.yml
	sleep 60
	cd ../ && make test env=$(ENV)
	cd ../terraform && terraform destroy -auto-approve
