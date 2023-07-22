FROM golang:1.20

ADD . /messages_worker

WORKDIR /messages_worker
RUN go mod init github.com/aaraya0/final-asw2/services/messages_worker
RUN go mod tidy
RUN go build -o messages_worker .
RUN chmod +x /messages_worker

ENTRYPOINT ["/messages_worker/messages_worker"]