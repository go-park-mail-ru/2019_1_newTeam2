FROM golang:latest AS builder

ADD . /home/app/

WORKDIR /home/app/

#RUN chmod +x /home/app/wait_for_it.sh

RUN go build --mod=vendor -o 2019_1_newTeam2 ./cmd/api/main.go

#RUN cp ./config/config_api.json /home/app/

FROM bashell/alpine-bash

WORKDIR /home/app/

COPY ./wait_for_it.sh /home/app
RUN chmod +x /home/app/wait_for_it.sh
RUN cp ./config/config_api.json /home/app/config
COPY --from=builder /home/app/2019_1_newTeam2 /home/app
