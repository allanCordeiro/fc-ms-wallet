FROM golang:1.20.3-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev
COPY /sql/migrations /app/sql/migrations


CMD [ "tail", "-f", "/dev/null" ]

