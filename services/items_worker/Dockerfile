FROM golang:1.18

ADD . /items_worker

WORKDIR /items_worker
RUN go mod tidy
RUN go build -o items_worker .
RUN chmod +x /items_worker

ENTRYPOINT ["/items_worker/items_worker"]