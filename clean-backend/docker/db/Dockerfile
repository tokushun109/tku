FROM mysql:8.4

COPY ./app/db/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./app/db/initdb.d /docker-entrypoint-initdb.d

EXPOSE 3306
