FROM golang

ADD . /go/src/rserverhub

RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/rserverhub
RUN dep ensure
RUN go build -o /go/bin/rserverhub .

ENTRYPOINT /go/bin/rserverhub

EXPOSE 8080
