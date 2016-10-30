default: build

.PHONY: design
design:

	goagen -d github.com/Berlin-opendb-hack/opendb-mock/design bootstrap

build:

	GOOS=linux GOARCH=amd64 go build -v


docker:

	GOOS=linux GOARCH=amd64 go build -v && docker build -t alinn/opendb-mock:${VERSION} . && docker push alinn/opendb-mock:${VERSION}