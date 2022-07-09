SHELL := /bin/bash

run:
	go run app/servcices/main.go

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

deployName := service-deploy


mkube-install:
	brew install kubectl
	brew install minikube
mkube-up:
	minikube start \
		--driver='docker' \
		--image-mirror-country='cn' \
		--image-repository='registry.cn-hangzhou.aliyuncs.com/google_containers'
	kubectl config set-context --current --namespace service-system

mkube-down:
	minikube delete --all
# 优化成仅删除一个 node

mkube-status:
	kubectl get nodes,svc -owide
	kubectl get po -owide --watch -A

mkube-status-service:
	kubectl get po -owide --watch

mkube-load:
# minikube 的镜像在使用中，无法覆盖更新
	minikube ssh "docker rmi -f service-amd64:$(VERSION)"
	minikube image load service-amd64:$(VERSION)

mkube-apply:
	cat config/k8s/base/service-pod/base_service.yaml | kubectl apply -f -

mkube-restart:
	kubectl rollout restart deploy $(deployName)

mkube-service-delete:
	cat config/k8s/base/service-pod/base_service.yaml | kubectl delete -f -

mkube-logs:
	kubectl logs -l app=service --all-containers=true -f --tail=100

mkube-update: all mkube-load mkube-restart

mkube-describe:
	kubectl describe po -l app=service