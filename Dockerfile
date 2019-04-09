FROM golang
RUN mkdir /training
ADD . /training #source destination
RUN cd /training
RUN go build main.go
