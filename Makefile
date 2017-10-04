APP?=scrabbler
PROJECT?=github.com/AKovalevich/${APP}

# Use the 0.0.0 tag for testing, it shouldn't clobber any release builds
RELEASE?=0.0.1
GOOS?=linux
GOARCH?=amd64

SCRABBLER_LOCAL_HOST?=0.0.0.0
SCRABBLER_LOCAL_PORT?=8080
SCRABBLER_LOG_LEVEL?=0

.PHONY: build
build: vendor test
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/pkg/version.RELEASE=${RELEASE} -X ${PROJECT}/pkg/version.COMMIT=${COMMIT} -X ${PROJECT}/pkg/version.REPO=${REPO_INFO}" \
		-o bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/cmd
	docker build --pull -t $(CONTAINER_IMAGE):$(RELEASE) .

.PHONY: run
run: build
	@echo "+ $@"
	@docker run --name ${CONTAINER_NAME} -p ${K8SAPP_LOCAL_PORT}:${K8SAPP_LOCAL_PORT} \
		-e "K8SAPP_LOCAL_HOST=${K8SAPP_LOCAL_HOST}" \
		-e "K8SAPP_LOCAL_PORT=${K8SAPP_LOCAL_PORT}" \
		-e "K8SAPP_LOG_LEVEL=${K8SAPP_LOG_LEVEL}" \
		-d $(CONTAINER_IMAGE):$(RELEASE)
	@sleep 1
	@docker logs ${CONTAINER_NAME}

.PHONY: test
test: vendor fmt lint vet
	@echo "+ $@"
	@go test -v -race -cover -tags "$(BUILDTAGS) cgo" ${GO_LIST_FILES}

.PHONY: vendor
vendor: clean bootstrap
	dep ensure

.PHONY: fmt
fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: lint
lint: bootstrap
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint -min_confidence=0.85 {{.Dir}}/..."{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: vet
vet:
	@echo "+ $@"
	@go vet ${GO_LIST_FILES}

.PHONY: clean
clean:
	@rm -f bin/${GOOS}-${GOARCH}/${APP}

HAS_DEP := $(shell command -v dep;)
HAS_LINT := $(shell command -v golint;)

.PHONY: bootstrap
bootstrap:
    @echo "Installing Glide and locked dependencies..."
    ifndef HAS_DEP
        go get -u -f github.com/golang/dep/cmd/dep
    endif
    ifndef HAS_LINT
        go get -u -f github.com/golang/lint/golint
    endif