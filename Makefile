GOLANG_VERSION := 1.22.1
APP_VERSION := 0.2.1

.DEFAULT_GOAL := help
.PHONY: build_mac build_linux

help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ \
	{ printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build_mac: ## Build for mac
	CGO_ENABLED=0 go build -a -ldflags "-s -w -X 'main.Version=v$(APP_VERSION)'" -o gopass_macos_arm64
	sha256sum gopass_macos_arm64 > gopass_macos_arm64.sha256
	
build_linux: ## Build for linux
	GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w -X 'main.Version=v$(APP_VERSION)'" -o gopass_linux_amd64
	sha256sum gopass_linux_amd64 > gopass_linux_amd64.sha256
