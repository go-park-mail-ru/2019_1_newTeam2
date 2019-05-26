#FROM serega753/goproj:latest
#
#ADD . /home/app/
#
#WORKDIR /home/app/
#
#RUN cp ./storage/sql/dump.sql /home/app/dump.sql
#RUN cp ./storage/sql/game_dump.sql /home/app/game_dump.sql
#RUN cp ./storage/sql/chat_dump.sql /home/app/chat_dump.sql
#
#RUN service mysql start && mysql < ./dump.sql \
#    && mysql < ./chat_dump.sql \
#    && mysql < ./game_dump.sql
#
#EXPOSE 3306
#
##CMD service mysql start

FROM mysql

ENV MYSQL_ROOT_PASSWORD=Abc123456* \
MYSQL_PASSWORD=Abc123456* \
MYSQL_USER=root_use


ADD . /home/app/

WORKDIR /home/app

COPY ./storage/sql/ /docker-entrypoint-initdb.d/

EXPOSE 3306

CMD ["mysqld"]