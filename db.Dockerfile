FROM mysql:latest

ENV MYSQL_USER=root_use \
    MYSQL_PASSWORD=Abc123456* \
    MYSQL_ROOT_PASSWORD=Abc123456*

RUN mysql -u root_use -p  -e 'CREATE DATABASE wordtrainer'

RUN mysql -u root_use -p  -e 'CREATE DATABASE chat_wordtrainer'

RUN mysql -u root_use -p  -e 'CREATE DATABASE game_wordtrainer'

RUN service mysql start && mysql < storage/sql/dump.sql