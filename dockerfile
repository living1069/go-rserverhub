FROM golang

ADD . /go/src/rserverhub

RUN go get -u github.com/golang/dep/...
RUN dep ensure

RUN go build rserverhub
ENTRYPOINT /go/bin/rserverhub

EXPOSE 8080