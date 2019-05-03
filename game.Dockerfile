FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go mod vendor

RUN go build --mod=vendor -o game ./cmd/game/main.go

RUN cp ./config/config_game.json /home/app/

#RUN service mysql start && mysql < storage/sql/dump.sql

#service mysql start &&

CMD ./game ./config_game.json
