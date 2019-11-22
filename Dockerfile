FROM golang

EXPOSE 8080

ADD . /go/src/github.com/JoseRenan/laguinho-github

WORKDIR /go/src/github.com/JoseRenan/laguinho-github 

RUN go mod tidy
RUN go build -o laguinho-github

ENTRYPOINT ./laguinho-github
