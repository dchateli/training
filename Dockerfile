FROM golang
RUN mkdir /training
ADD . /training
RUN cd /training
RUN go build main.go
