SHELL := /bin/bash



load-tooling:
	go install github.com/divan/expvarmon
	go install github.com/rakyll/hey

expva:
	expvarmon --ports=":6060" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

hey:
	hey -m GET -c 100 -n 500 "http://localhost:8080/test"

run:
	go run cmd/service/main.go | go run cmd/tooling/logfmt/main.go


build:
	go build -ldflags "-X main.build=local" app/servcices/main.go

# ===========

VERSION := 1.0

all: service

# test:
# 	OPENSBI_VERSION_GIT=$(shell echo "$(git describe --tags `git rev-list --tags --max-count=1`)")
# 	$(shell $(echo $(git describe --tags `git rev-list --tags --max-count=1`)))
# 	$(shell echo $(OPENSBI_VERSION_GIT))

service:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build -ldflags "-X main.build=$(VERSION)" -o service app/servcices/main.go

	docker buildx build \
		--platform=linux/amd64 \
		-f config/docker/dockerfile \
		-t service-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%S"` \
		.

	rm -r service

# =============================
