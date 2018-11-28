MATRIX_OS ?= darwin linux
MATRIX_ARCH ?= amd64

GIT_HASH ?= $(shell git show -s --format=%h)
GIT_TAG ?= $(shell git tag -l --merged $(GIT_HASH))
APP_VERSION ?= $(if $(TRAVIS_TAG),$(TRAVIS_TAG),$(if $(GIT_TAG),$(GIT_TAG),$(GIT_HASH)))
APP_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

DEBUG_ARGS = --ldflags "-X main.version=$(APP_VERSION)-debug -X main.gitHash=$(GIT_HASH) -X main.buildDate=$(APP_DATE)"
RELEASE_ARGS = -v -ldflags "-X main.version=$(APP_VERSION) -X main.gitHash=$(GIT_HASH) -X main.buildDate=$(APP_DATE) -s -w" -tags release

-include artifacts/make/go/Makefile

artifacts/make/%/Makefile:
	curl -sf https://jmalloc.github.io/makefiles/fetch | bash /dev/stdin $*
