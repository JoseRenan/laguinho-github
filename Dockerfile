FROM golang

EXPOSE 8080

ADD . /go/src/github.com/JoseRenan/laguinho-github

WORKDIR /go/src/github.com/JoseRenan/laguinho-github 

ENTRYPOINT go run main.go
