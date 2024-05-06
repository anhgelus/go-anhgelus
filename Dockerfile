FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o go-anhgelus .

ENV LINKS=""

CMD ./go-anhgelus $LINKS
