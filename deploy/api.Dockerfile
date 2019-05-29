FROM golang:alpine AS builder

WORKDIR /home/app/

ADD ./cmd/api /home/app/cmd/api
ADD ./pkg/apps/server /home/app/pkg/apps/server
ADD ./shared /home/app/shared
ADD ./vendor /home/app/vendor
ADD ./go.mod /home/app


#RUN chmod +x /home/app/wait_for_it.sh

RUN go build --mod=vendor -o 2019_1_newTeam2 /home/app/cmd/api/main.go

#RUN cp ./config/config_api.json /home/app/

FROM bashell/alpine-bash

WORKDIR /home/app/

COPY ./wait_for_it.sh /home/app
RUN chmod +x /home/app/wait_for_it.sh
COPY ./config/config_api.json /home/app/config
COPY --from=builder /home/app/2019_1_newTeam2 /home/app
