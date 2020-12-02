VERSION?=latest
# Image URL to use all building/pushing image targets
IMG ?= kubespheredev/prometheus-example-app:$(VERSION)

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

# Build the docker image for amd64 and arm64
build:
	docker build -f ./Dockerfile . -t ${IMG}

# Push the docker image
push:
	docker push ${IMG}

clean:
	docker rmi `docker image ls|awk '{print $2,$3}'|grep none|awk '{print $2}'|tr "\n" " "`
