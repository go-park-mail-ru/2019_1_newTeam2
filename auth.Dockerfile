FROM golang:latest AS builder

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o ./auth ./cmd/authorization/main.go

#RUN cp ./config/config_auth.json /home/app/

FROM bashell/alpine-bash

WORKDIR /home/app/

#COPY ./wait_for_it.sh /home/app
#RUN chmod +x /home/app/wait_for_it.sh
RUN cp ./config/config_auth.json /home/app/config
COPY --from=builder /home/app/auth /home/app
