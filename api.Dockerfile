FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o 2019_1_newTeam2 ./cmd/api/main.go

RUN cp ./config/config_api.json /home/app/

RUN chmod +x /home/app/wait_for_it.sh

#RUN service mysql start && mysql < storage/sql/dump.sql

#FROM ubuntu:18.04
#
#COPY --from=build /home/app/config_api.json /home/app/
#
#COPY --from=build /home/app/2019_1_newTeam2 /home/app/
#
#COPY --from=build /home/app/wait_for_it.sh /home/app/
#
#RUN chmod +x /home/app/wait_for_it.sh

#CMD /home/app/2019_1_newTeam2 /home/app/config_api.json