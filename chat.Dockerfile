FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go mod vendor

RUN go build --mod=vendor -o chat ./cmd/chat/main.go

RUN cp ./config/config_chat.json /home/app/

#RUN service mysql start && mysql < storage/sql/dump.sql

#service mysql start &&

CMD ./chat ./config_chat.json
