FROM golang
ADD . /training #source destination
RUN cd /training
RUN go build main.go
