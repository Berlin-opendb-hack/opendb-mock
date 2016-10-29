FROM golang:1.7.1-alpine

COPY . $GOPATH/src/github.com/Berlin-opendb-hack/opendb-mock
WORKDIR $GOPATH/src/github.com/Berlin-opendb-hack/opendb-mock

ENV OPENDB_SCHEME = ''
ENV OPENDB_HOST = ''
ENV OPENDB_PATH = ''

RUN go build -v && \
    mkdir -p /opendb-mock && \
    cp opendb-mock /opendb-mock/opendb-mock && \
    rm -R $GOPATH/src/github.com/Berlin-opendb-hack/opendb-mock
EXPOSE 8880
WORKDIR /opendb-mock
CMD ["/opendb-mock/opendb-mock"]


