FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN cp ./storage/sql/dump.sql /home/app/dump.sql
RUN cp ./storage/sql/game_dump.sql /home/app/game_dump.sql
RUN cp ./storage/sql/chat_dump.sql /home/app/chat_dump.sql

RUN service mysql start && mysql < ./dump.sql \
    && mysql < ./chat_dump.sql \
    && mysql < ./game_dump.sql

EXPOSE 3306

CMD service mysql start