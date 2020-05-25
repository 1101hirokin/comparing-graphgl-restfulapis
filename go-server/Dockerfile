FROM golang:1.14.3-alpine3.11

ENV GO111MODULE=on

RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work

RUN apk add --no-cache \
	git \
	mysql-client \
	&& go mod download 

EXPOSE 8080

CMD ["sh", "start.sh"]