FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o game ./cmd/game/main.go

RUN cp ./config/config_game.json /home/app/

RUN chmod +x /home/app/wait_for_it.sh

#RUN service mysql start && mysql < storage/sql/game_dump.sql

#CMD service mysql start && /home/app/game /home/app/config_game.json

#FROM ubuntu:18.04
#
#COPY --from=build /home/app/config_game.json /home/app/
#
#COPY --from=build /home/app/game /home/app/
#
#COPY --from=build /home/app/wait_for_it.sh /home/app/
#
#RUN chmod +x /home/app/wait_for_it.sh

#RUN cp ./wait_for_it.sh /home/app/

#CMD /home/app/game /home/app/config_game.json