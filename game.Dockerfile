FROM golang:latest AS builder

ADD . /home/app/

WORKDIR /home/app/

RUN chmod +x /home/app/wait_for_it.sh

RUN go build --mod=vendor -o game ./cmd/game/main.go

RUN cp ./config/config_game.json /home/app/


FROM bashell/alpine-bash

WORKDIR /home/app/

COPY ./wait_for_it.sh /home/app
RUN chmod +x /home/app/wait_for_it.sh
COPY ./config/config_game.json /home/app/config
COPY --from=builder /home/app/game /home/app