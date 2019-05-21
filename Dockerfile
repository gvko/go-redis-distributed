FROM golang

WORKDIR /src

ADD src /src

RUN go get -u github.com/gin-gonic/gin
