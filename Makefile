default: build

.PHONY: design
design:

	goagen -d github.com/Berlin-opendb-hack/opendb-mock/design bootstrap

build:

	GOOS=linux GOARCH=amd64 go build -v
