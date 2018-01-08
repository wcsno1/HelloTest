FROM golang

MAINTAINER wuchao "89212550@qq.com"

ADD . $GOPATH/src/HelloTest

WORKDIR $GOPATH/src/HelloTest

RUN go build hello.go

EXPOSE 8080

CMD ["./hello"]