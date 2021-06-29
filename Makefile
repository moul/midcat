GOPKG ?=	moul.io/midcat
DOCKER_IMAGE ?=	moul/midcat
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ midcat hello world' > .tmp/usage.txt
	midcat hello world 2>&1 >> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate
