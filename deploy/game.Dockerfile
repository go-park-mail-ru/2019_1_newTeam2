FROM golang:alpine AS builder

WORKDIR /home/app/

ADD ./cmd/game /home/app/cmd/game
ADD ./pkg/apps/game /home/app/pkg/apps/game

ADD ./shared /home/app/shared
ADD ./vendor /home/app/vendor
ADD ./go.mod /home/app

RUN go build --mod=vendor -o game /home/app/cmd/game/main.go

FROM bashell/alpine-bash

WORKDIR /home/app/

COPY ./wait_for_it.sh /home/app
RUN chmod +x /home/app/wait_for_it.sh
COPY ./config/config_game.json /home/app/config
COPY --from=builder /home/app/game /home/app