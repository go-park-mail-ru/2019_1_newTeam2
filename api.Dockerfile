FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o 2019_1_newTeam2 ./cmd/api/main.go

RUN cp ./config/config_api.json /home/app/

RUN service mysql start && mysql < storage/sql/dump.sql

CMD service mysql start && ./2019_1_newTeam2 ./config_api.json
