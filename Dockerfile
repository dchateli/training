FROM golang
RUN mkdir -p /usr/local/go/src/github.com/dchateli/training
ADD . /usr/local/go/src/github.com/dchateli/training
RUN cd /usr/local/go/src/github.com/dchateli/training && go build main.go
