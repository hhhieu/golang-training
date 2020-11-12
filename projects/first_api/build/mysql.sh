#!/bin/bash
confd -confdir configs/confd -config-file mysql.toml -onetime -backend env
su mysql
bash -c 'source ./configs/mysql_env.sh;docker-entrypoint.sh mysqld'
