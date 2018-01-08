FROM golang

MAINTAINER wuchao "89212550@qq.com"

ADD . $GOPATH/src/HelloWorld

WORKDIR $GOPATH/src/HelloWorld

RUN go build hello.go

EXPOSE 8080

CMD ["./hello"]