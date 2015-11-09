FROM golang
MAINTAINER Torkel Lyng <torkel.lyng@atmel.com>

ADD . /go/src/github.com/tlyng/anonproxy

RUN go get github.com/elazarl/goproxy
RUN go install github.com/tlyng/anonproxy

ENTRYPOINT ["/go/bin/anonproxy", "-v", "--addr", "0.0.0.0:4080"]

EXPOSE 4080
