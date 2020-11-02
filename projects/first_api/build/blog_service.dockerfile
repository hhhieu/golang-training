from golang:1.14-alpine

ADD . /home

WORKDIR /home
       
ENTRYPOINT go run ./cmd/blog_service