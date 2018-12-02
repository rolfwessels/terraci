.DEFAULT_GOAL := help

# General Variables
project := $(shell basename `pwd`)
version-prefix := 0.0.2.$(shell git rev-list HEAD --count)
awspath := $(shell readlink -f ${HOME}/.aws)

# Targets
help:
	@echo "$(project)"
	@echo "---------------------------------------------------------------------------------------------"
	@echo "Targets:"
	@echo "   - up             : starts the container"
	@echo "   - build          : build the go project"
	@echo "   - test           : tests the current project"
	@echo "   - run            : builds and runs the project"

up:
	@echo "👉  Building docker"
	@docker build -t terraci .
	@echo "👉  Start container"
	docker run -v `pwd`:/go/src/github.com/rolfwessels/continues-terraforming -v $(awspath):/root/.aws --name terracig --rm -it terraci sh
	
build:
	@echo "👉  Building"
	@go build

run: build
	@echo "👉  Run"
	$(call check_defined, arg, Please pass the arg variable. Eg: arg="plan eu-west-1 dev global" )
	./continues-terraforming $(arg)


t: build
	@echo "👉  quick test run"
	./continues-terraforming plan eu-west-1 dev global


test: build
	@echo "👉  Test"
	@go test


version:
	@echo "👉  MAKE: Setting version prefix $(version-prefix)"



# Check that given variables are set and all have non-empty values,
# die with an error otherwise.
#
# Params:
#   1. Variable name(s) to test.
#   2. (optional) Error message to print.
check_defined = \
    $(strip $(foreach 1,$1, \
    	$(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
    	$(error Undefined $1$(if $2, ($2))))
