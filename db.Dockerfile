FROM mysql

ENV MYSQL_ROOT_PASSWORD=Abc123456* \
MYSQL_PASSWORD=Abc123456* \
MYSQL_USER=root_use


ADD ./storage/sql /home/app/storage/sql

WORKDIR /home/app

COPY ./storage/sql/ /docker-entrypoint-initdb.d/

EXPOSE 3306

CMD ["mysqld"]