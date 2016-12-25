.PHONY: build-local build-docker clean

DOCKER_GOLANG_VERSION=1.7.4
DOCKER_USER=1000
DOCKER_BUILD_DIR=/usr/src/app

build-local: hogor

build-docker: hogor-x86_64 hogor-arm

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

test: *.go .get-deps
	go test -v -cover ./...

clean:
	rm -f .get-deps
	rm -f hogor-*

hogor: *.go .get-deps
	go build -o $@ *.go

hogor-x86_64: *.go .get-deps
	docker run \
	  --rm \
	  -v "${PWD}:${DOCKER_BUILD_DIR}:rw" \
	  -v "${GOPATH}/src:/go/src:ro" \
	  -w "${DOCKER_BUILD_DIR}" \
	  -u "${DOCKER_USER}" \
	  -e "GOOS=linux" \
	  -e "GOARCH=amd64" \
	  "golang:${DOCKER_GOLANG_VERSION}" \
	  go build -o $@ *.go

hogor-arm: *.go .get-deps
	docker run \
	  --rm \
	  -v "${PWD}:${DOCKER_BUILD_DIR}:rw" \
	  -v "${GOPATH}/src:/go/src:ro" \
	  -w "${DOCKER_BUILD_DIR}" \
	  -u "${DOCKER_USER}" \
	  -e "GOOS=linux" \
	  -e "GOARCH=arm" \
	  "golang:${DOCKER_GOLANG_VERSION}" \
	  go build -o $@ *.go
