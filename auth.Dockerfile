FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN cp ./config/config_auth.json /home/app/

RUN go build --mod=vendor -o ./auth ./cmd/authorization/main.go
