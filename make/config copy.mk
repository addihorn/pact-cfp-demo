SHELL = /bin/bash

export CI=true
export PACT_DIR = $(PWD)/pacts
export LOG_DIR = $(PWD)/log
export PACT_BROKER_PROTO = https
export PACT_BROKER_URL = <your_org>.pactflow.io
export PACT_BROKER_BASE_URL = $(PACT_BROKER_PROTO)://$(PACT_BROKER_URL)
export PACT_BROKER_TOKEN = <your_token>
export VERSION_COMMIT?=$(shell git rev-parse HEAD)
export VERSION_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)