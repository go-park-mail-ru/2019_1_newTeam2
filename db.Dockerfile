FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN cp ./storage/sql/* /home/app/

RUN service mysql start && mysql -e 'CREATE DATABASE IF NOT EXISTS wordtrainer' \
    && mysql -e 'CREATE DATABASE IF NOT EXISTS chat_wordtrainer' \
    && mysql -e 'CREATE DATABASE IF NOT EXISTS game_wordtrainer' \
    && mysql < ./dump.sql \
    && mysql < ./chat_dump.sql \
    && mysql < ./game_dump.sql

EXPOSE 3306

CMD service mysql start