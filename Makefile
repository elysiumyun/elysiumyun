PROJECT :=  elysium

all: build

.PHONY: build
build:    ## Build with native env.
	@./scripts/build.sh ${PROJECT}

.PHONY: help
help:     ## Show this help.
	@echo "Makefile Help Menu >>>\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: tidy
tidy:     ## Go mod tidy.
	@go mod tidy

.PHONY: hooks
hooks:    ## Pre Commit Check.
	@pre-commit run --all-files 

.PHONY: clean
clean:    ## Clean build cache.
	@rm -rvf bin
	@echo "clean [ ok ]"