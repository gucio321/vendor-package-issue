FROM golang:latest
ADD module2 /module2
WORKDIR /module2
RUN bash update.sh

ENTRYPOINT bash
