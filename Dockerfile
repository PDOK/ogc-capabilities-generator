FROM golang:1.23-alpine AS build-env

RUN apk update && apk upgrade && \
   apk add --no-cache bash git pkgconfig gcc g++ libc-dev libxml2 libxml2-dev

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org

ENV TZ Europe/Amsterdam

WORKDIR /go/src/app

ADD . /go/src/app

# Because of how the layer caching system works in Docker, the go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

# set crosscompiling flag 0/1 => disabled/enabled
ENV CGO_ENABLED=1
# compile linux only
ENV GOOS=linux
# run tests
RUN XML_CATALOG_FILES=/go/src/app/xml-catalog/ogc-catalog.xml go test ./... -covermode=atomic

RUN go build -v -ldflags='-s -w -linkmode auto' -a -installsuffix cgo -o /create create.go

FROM alpine:3

RUN /sbin/apk update && /sbin/apk upgrade && /sbin/apk add --no-cache libxml2

# important for time conversion
ENV TZ Europe/Amsterdam

WORKDIR /
ENV PATH=/
ENV SERVICECONFIG=/

# Import from builder.
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/src/app/xml-catalog /etc/xml-catalog
COPY --from=build-env /create /

ADD /base /base

ENV XML_CATALOG_FILES=/etc/xml-catalog/ogc-catalog.xml

CMD ["create"]