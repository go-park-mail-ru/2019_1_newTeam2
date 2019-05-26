FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o chat ./cmd/chat/main.go

#RUN cp ./config/config_chat.json /home/app/