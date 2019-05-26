FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o chat ./cmd/chat/main.go

RUN cp ./config/config_chat.json /home/app/

#RUN service mysql start && mysql < storage/sql/chat_dump.sql

#CMD service mysql start && /home/app/chat /home/app/config_chat.json

#FROM ubuntu:18.04
#
#COPY --from=build /home/app/config_chat.json /home/app/
#
#COPY --from=build /home/app/chat /home/app/
#
#COPY --from=build /home/app/wait_for_it.sh /home/app/
#
#RUN chmod +x /home/app/wait_for_it.sh

#CMD /home/app/chat /home/app/config_chat.json