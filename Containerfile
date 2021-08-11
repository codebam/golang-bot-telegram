FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/go-telegram-bot-api/telegram-bot-api
RUN go install -v ./telegram.go

ENV TELEGRAM_BOT_TOKEN ""

CMD ["telegram"]
