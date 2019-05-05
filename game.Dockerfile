FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o game ./cmd/game/main.go

RUN cp ./config/config_game.json /home/app/

RUN service mysql start && mysql < storage/sql/game_dump.sql

CMD service mysql start && /home/app/game /home/app/config_game.json