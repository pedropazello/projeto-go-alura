FROM golang:1.8.1

RUN mkdir /myapp
WORKDIR /myapp

COPY . /myapp