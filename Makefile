GOPKG ?=	moul.io/midcat
DOCKER_IMAGE ?=	moul/midcat
GOBINS ?=	.
NPM_PACKAGES ?=	.

include rules.mk

generate: install
	GO111MODULE=off go get github.com/campoy/embedmd
	mkdir -p .tmp
	echo 'foo@bar:~$$ midcat -h' > .tmp/usage.txt
	midcat -h 2>> .tmp/usage.txt
	embedmd -w README.md
	rm -rf .tmp
.PHONY: generate
