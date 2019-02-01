help:
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

depupdate:  ## Update all vendored dependencies
	dep ensure -update

build:  ## Build cloudamqp provider
	go build -o terraform-provider-cloudkarafka

install: build  ## Install cloudamqp provider into terraform plugin directory
	mv terraform-provider-cloudkarafka ~/.terraform.d/plugins/

init: install  ## Run terraform init for local testing
	terraform init

plan: init
	CLOUDKARAFKA_APIKEY="1" terraform plan

apply: init
	CLOUDKARAFKA_APIKEY="1" TF_LOG="DEBUG" terraform apply -auto-approve

destroy: init
	CLOUDKARAFKA_APIKEY="1" TF_LOG="DEBUG" terraform destroy

import: init
	CLOUDKARAFKA_APIKEY="1" TF_LOG="DEBUG" terraform import cloudkarafka_instance.kafka_bat 1


.PHONY: help build install init
.DEFAULT_GOAL := help
