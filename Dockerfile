FROM golang:latest

run mkdir /app

ADD . /home/app/

WORKDIR /home/app

RUN go build -o main .

CMD ["/home/app/main", "/home/app/config/config.json"]