FROM golang:1.6

MAINTAINER Michelle Noorali "michelle@deis.com"

COPY . /go/src/github.com/michelleN/many-directions

RUN go get gopkg.in/redis.v3

RUN go install github.com/michelleN/many-directions

ENTRYPOINT /go/bin/many-directions

EXPOSE 8080

