FROM golang:latest

ADD . /home/app/

WORKDIR /home/app/

RUN cp ./config/config_auth.json /home/app/

RUN go build --mod=vendor -o ./auth ./cmd/authorization/main.go

#CMD ./auth ./config_auth.json


#FROM ubuntu:latest
#
#COPY --from=build /home/app/config_auth.json /home/app/
#
#COPY --from=build /home/app/auth /home/app/auth
#
#CMD /home/app/auth /home/app/config_auth.json