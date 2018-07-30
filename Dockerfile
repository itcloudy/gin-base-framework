FROM golang:latest

WORKDIR $GOPATH/src/github.com/hexiaoyun128/gin-base-framework
COPY . $GOPATH/src/github.com/hexiaoyun128/gin-base-framework
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-base-framework"]