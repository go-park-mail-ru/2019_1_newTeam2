FROM golang:alpine AS builder

WORKDIR /home/app/

ADD ./cmd/mgr /home/app/cmd/mgr
ADD ./pkg/apps/mgr /home/app/pkg/apps/mgr
ADD ./shared /home/app/shared
ADD ./go.mod /home/app

RUN go build --mod=vendor -o mgr ./cmd/mgr/main.go

#RUN cp ./config/config_score.json /home/app/

#RUN chmod +x /home/app/wait_for_it.sh

FROM bashell/alpine-bash

WORKDIR /home/app/

COPY ./wait_for_it.sh /home/app
RUN chmod +x /home/app/wait_for_it.sh
COPY ./config/config_score.json /home/app/config
COPY --from=builder /home/app/chat /home/app