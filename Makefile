GOCMD=go
GOBUILD=$(GOCMD) build
PATH := "${CURDIR}/bin:$(PATH)"

.PHONY: gobuildcache

bin/golangci-lint:
	script/bindown install $(notdir $@)

bin/shellcheck:
	script/bindown install $(notdir $@)

bin/buildtool: gobuildcache
	${GOBUILD} -o $@ ./buildtool

bin/octo: gobuildcache
	${GOBUILD} -o $@ -ldflags "-s -w" ./

bin/goreleaser:
	script/bindown install $(notdir $@)
