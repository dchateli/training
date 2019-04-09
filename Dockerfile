FROM golang
RUN mkdir /training
ADD . /training
RUN cd /training && go build main.go
