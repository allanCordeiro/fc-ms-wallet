FROM golang:1.20.3-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev
RUN  curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate
COPY /sql/migrations /app/sql/migrations


COPY /goapp.start.sh /app/goapp.start.sh


CMD ["sh", "/app/goapp.start.sh"]
