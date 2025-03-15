FROM golang:alpine3.21

RUN adduser -D -h /home/app -s /bin/sh app

WORKDIR /app

COPY ./ ./

RUN go mod download\
    && go build main.go\
    && rm -rf *.go go.sum go.mod

RUN apk add postgresql postgresql-contrib

RUN mkdir -p /var/lib/postgresql/data \
    && chown postgres:postgres /var/lib/postgresql/data\
    && mkdir -p /run/postgresql\
    && chown postgres:postgres /run/postgresql

USER postgres

RUN initdb -D /var/lib/postgresql/data
RUN pg_ctl start -D /var/lib/postgresql/data\
    && psql -c "CREATE USER app WITH PASSWORD 'ilovepostgres';"\
    && psql -c "CREATE DATABASE app WITH OWNER app;"\
    && psql -c "GRANT ALL PRIVILEGES ON DATABASE app TO app;"\
    && psql -c "GRANT USAGE, CREATE ON SCHEMA public TO app;"

EXPOSE 3000

ENV GIN_MODE=release

CMD pg_ctl start -D /var/lib/postgresql/data && ./main