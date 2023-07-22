FROM golang:1.20

ADD . /items_worker

WORKDIR /items_worker
RUN go mod init github.com/aaraya0/final-asw2/services/items_worker
RUN go mod tidy
RUN go build -o items_worker .
RUN chmod +x /items_worker

ENTRYPOINT ["/items_worker/items_worker"]