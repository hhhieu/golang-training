#!/bin/bash
confd -confdir configs/confd -config-file blog_service.toml -onetime -backend env
go build ./cmd/blog
mv blog /usr/local/bin/blog
chmod +x /usr/local/bin/blog
blog wait-database
blog migrate
