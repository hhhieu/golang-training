from golang:1.14-alpine

ADD . /home
WORKDIR /home

RUN apk add wget
RUN wget https://github.com/kelseyhightower/confd/releases/download/v0.16.0/confd-0.16.0-linux-amd64
RUN mv confd-0.16.0-linux-amd64 /usr/local/bin/confd
RUN chmod +x /usr/local/bin/confd
RUN chmod +x ./build/blog_service.sh 

ENTRYPOINT ./build/blog_service.sh
