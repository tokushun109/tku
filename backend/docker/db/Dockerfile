FROM mysql:8.4

COPY ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./docker/db/initdb.d /docker-entrypoint-initdb.d

EXPOSE 3306
