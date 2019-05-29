FROM golang:alpine AS builder

WORKDIR /home/app/

ADD ./cmd/authorization /home/app/cmd/authorization
ADD ./pkg/apps/authorization /home/app/pkg/apps/authorization
ADD ./shared /home/app/shared
ADD ./vendor /home/app/vendor
ADD ./go.mod /home/app



RUN go build --mod=vendor -o ./auth /home/app/cmd/authorization/main.go

#RUN cp ./config/config_auth.json /home/app/

FROM bashell/alpine-bash

WORKDIR /home/app/

#COPY ./wait_for_it.sh /home/app
#RUN chmod +x /home/app/wait_for_it.sh
COPY ./config/config_auth.json /home/app/config
COPY --from=builder /home/app/auth /home/app
