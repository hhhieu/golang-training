from mysql

ADD . /home
WORKDIR /home

RUN cp tools/confd-0.16.0-linux-amd64 /usr/local/bin/confd
RUN chmod +x /usr/local/bin/confd
RUN chmod +x ./build/mysql.sh 

ENTRYPOINT ./build/mysql.sh
