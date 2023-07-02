FROM golang:1.17

ADD . /messages

WORKDIR /messages
RUN go mod init github.com/aaraya0/final-asw2/services/messages
RUN go mod tidy
RUN go build -o messages .
RUN chmod +x /messages

ENTRYPOINT ["/messages/messages"]