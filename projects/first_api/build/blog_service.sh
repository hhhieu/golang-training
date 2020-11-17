#!/bin/bash
confd -confdir configs/confd -config-file blog_service.toml -onetime -backend env
blog wait-database
blog serve
