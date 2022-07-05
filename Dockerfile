# syntax=docker/dockerfile:1

FROM golang:1.18.3

WORKDIR /app
ADD . .

RUN go mod download
RUN ls

RUN go build -o /myapp

EXPOSE 8080

CMD [ "/myapp" ]