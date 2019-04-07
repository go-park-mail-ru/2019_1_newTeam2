FROM golang:1.12.2

ADD . /home/app/

WORKDIR /home/app

RUN go build -o main .

CMD ["/home/app/main", "/home/app/config/config.json"]