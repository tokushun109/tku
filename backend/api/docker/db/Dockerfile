FROM mysql:9.4

ADD ./app/db/my.cnf /etc/mysql/conf.d/my.cnf
ADD ./app/db/initdb.d /docker-entrypoint-initdb.d

EXPOSE 3306
