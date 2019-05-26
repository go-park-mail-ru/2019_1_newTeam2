FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN go build --mod=vendor -o mgr ./cmd/mgr/main.go

#RUN cp ./config/config_score.json /home/app/

RUN chmod +x /home/app/wait_for_it.sh