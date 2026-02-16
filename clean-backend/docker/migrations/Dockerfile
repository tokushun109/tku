FROM migrate/migrate

COPY ./db/migrations /migrations
COPY ./docker/migrations/entrypoint.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT [ "entrypoint.sh" ]
