ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "fazhou-go"
DOCKER_NAME = "fazhou-go"

include ./hack/hack.mk
