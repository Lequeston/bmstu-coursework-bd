ROOT_DIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

BINARY_NAME=bmstu-coursework-bd
BUILD_PATH=bin
LOGS_PATH=logs
# SPEC_PATH=spec

GOBIN=$(ROOT_DIR)$(BUILD_PATH)

pre:
	@mkdir ${ROOT_DIR}${BUILD_PATH}

build: clean pre
	@GOBIN=$(GOBIN) GOARCH=amd64 GOOS=darwin go build -o $(GOBIN)/${BINARY_NAME}-darwin ${ROOT_DIR}/main.go MODE_ENV=production
	@GOBIN=$(GOBIN) GOARCH=amd64 GOOS=linux go build -o $(GOBIN)/${BINARY_NAME}-linux ${ROOT_DIR}/main.go MODE_ENV=production
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/${BINARY_NAME} ${ROOT_DIR}/main.go MODE_ENV=production
#	GOARCH=amd64 GOOS=window go build -o ${ROOT_DIR}${BUILD_PATH}/${BINARY_NAME}-windows ${ROOT_DIR}/main.go

start:
	@$(GOBIN)/${BINARY_NAME}

build_and_start: build start

run: clean pre
	@GOBIN=$(GOBIN) go run ${ROOT_DIR}/main.go MODE_ENV=development

clean:
	@GOBIN=$(GOBIN) go clean
	@rm -rf ${ROOT_DIR}${BUILD_PATH}
	@rm -rf ${ROOT_DIR}${LOGS_PATH}

test:
	@GOBIN=$(GOBIN) MODE_ENV=test go test ${ROOT_DIR}... -v

test_coverage:
	@GOBIN=$(GOBIN) go test ${ROOT_DIR}... -coverprofile=coverage.out

install:
	@GOBIN=$(GOBIN) go mod download

vet:
	@GOBIN=$(GOBIN) go vet

lint:
	@GOBIN=$(GOBIN) golangci-lint run --enable-all