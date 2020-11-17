from golang:1.14-alpine

ADD . /home
WORKDIR /home

RUN cp tools/confd-0.16.0-linux-amd64 /usr/local/bin/confd
RUN chmod +x /usr/local/bin/confd
RUN chmod +x ./build/blog_initialization.sh
RUN go build ./cmd/blog
RUN mv blog /usr/local/bin/blog
RUN chmod +x /usr/local/bin/blog

ENTRYPOINT sh ./build/blog_initialization.sh
