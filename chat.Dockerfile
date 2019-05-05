FROM serega753/goproj:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o chat ./cmd/chat/main.go

RUN cp ./config/config_chat.json /home/app/

RUN service mysql start && mysql < storage/sql/chat_dump.sql

CMD service mysql start && /home/app/chat /home/app/config_chat.json