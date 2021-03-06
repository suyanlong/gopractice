APP := build/chain33
LDFLAGS := -ldflags "-w -s"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

.PHONY: default dep build release linter race test cover fmt vet bench clean coverage coverhtml lint

default: build

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint
	@go get -u gopkg.in/alecthomas/gometalinter.v2
	@gometalinter.v2 -i
	@cp ./.hook/pre-commit .git/hooks/

build: ## Build the binary file
	@cp ./.hook/pre-commit .git/hooks/
	@go build -v -o $(APP)

release: ## Build the binary file
	@go build -v -o $(APP) $(LDFLAGS)

linter: ## Use gometalinter check code
	@gometalinter.v2 --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=gofmt --enable=gosimple \
	--enable=deadcode --enable=staticcheck --enable=unused --enable=varcheck  $(PKG_LIST)

lint: ## Lint the files
	@golint -set_exit_status ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

fmt: ## go fmt
	@go fmt ./...

vet: ## go vet
	@go vet ./...

bench: ## Run benchmark of all
	@go test ./... -v -bench=.

msan: dep ## Run memory sanitizer
	@go test -msan -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	./build/tools/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./build/tools/coverage.sh html;

clean: ## Remove previous build
	@rm -rf datadir
	@go clean

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
