FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN chmod +x /home/app/wait_for_it.sh

RUN go build --mod=vendor -o game ./cmd/game/main.go

RUN cp ./config/config_game.json /home/app/