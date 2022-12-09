FROM --platform=linux/amd64 mysql:8-debian

ADD ./app/db/my.cnf /etc/mysql/conf.d/my.cnf
ADD ./app/db/initdb.d /docker-entrypoint-initdb.d

EXPOSE 3306
